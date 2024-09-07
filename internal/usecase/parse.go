package usecase

import (
	"encoding/json"
	"fmt"
	"parse_server/internal/domain/repository"
	"parse_server/internal/domain/usecase"
	"time"
)

type EthereumParserParam struct {
	Storage      repository.Storage
	Notification usecase.Notification
	EthClient    repository.ETHClient
}

// EthereumParser 實現了 Parser interface
type EthereumParser struct {
	storage      repository.Storage
	notification usecase.Notification
	ethClient    repository.ETHClient
	currentBlock int
}

func NewEthereumParser(param EthereumParserParam) usecase.Parser {
	return &EthereumParser{
		storage:      param.Storage,
		notification: param.Notification,
		ethClient:    param.EthClient,
		currentBlock: 0,
	}
}

// fetchBlockTransactions 根據區塊號獲取區塊的交易
func (p *EthereumParser) fetchBlockTransactions(blockNumber string) ([]repository.Transaction, error) {
	result, err := p.ethClient.CallEthereum("eth_getBlockByNumber", []any{blockNumber, true})
	if err != nil {
		return nil, err
	}

	var rpcResponse repository.BlockResult
	err = json.Unmarshal(result, &rpcResponse)
	if err != nil {
		return nil, err
	}
	reply := make([]repository.Transaction, 0, len(rpcResponse.Result.Transactions))
	for _, item := range rpcResponse.Result.Transactions {
		to := ""
		if item.To != nil {
			to = *item.To
		}
		reply = append(reply, repository.Transaction{
			BlockHash:   rpcResponse.Result.Hash,
			BlockNumber: rpcResponse.Result.Number,
			From:        item.From,
			To:          to,
			Value:       item.Value,
		})
	}

	return reply, nil
}

// UpdateCurrentBlock 更新目前區塊
func (p *EthereumParser) UpdateCurrentBlock() error {
	result, err := p.ethClient.CallEthereum("eth_blockNumber", []any{})
	if err != nil {
		return err
	}

	var rpcResponse repository.EthereumRPCResponse
	err = json.Unmarshal(result, &rpcResponse)
	if err != nil {
		return err
	}

	// 將十六進制的區塊號轉為整數
	var blockNumber int
	_, err = fmt.Sscanf(rpcResponse.Result, "0x%x", &blockNumber)
	if err != nil {
		return err
	}
	p.currentBlock = blockNumber
	return nil
}

// GetCurrentBlock 取得當前區塊號
func (p *EthereumParser) GetCurrentBlock() int {
	return p.currentBlock
}

// Subscribe 訂閱地址
func (p *EthereumParser) Subscribe(address string) bool {
	p.storage.SubscribeAddress(address)
	return true
}

// GetTransactions 取得指定地址的交易
func (p *EthereumParser) GetTransactions(address string) []usecase.Transaction {
	r := p.storage.GetTransactions(address)
	result := make([]usecase.Transaction, 0, len(r))
	for _, item := range r {
		result = append(result, usecase.Transaction{
			BlockHash:   item.BlockHash,
			BlockNumber: item.BlockNumber,
			From:        item.From,
			To:          item.To,
			Value:       item.Value,
		})
	}

	return result
}

// FetchTransactionsForAddress 檢查與訂閱地址相關的交易並通知
func (p *EthereumParser) FetchTransactionsForAddress(address string) {
	blockNumber := fmt.Sprintf("0x%x", p.currentBlock)
	transactions, err := p.fetchBlockTransactions(blockNumber)
	if err != nil {
		fmt.Println("Error fetching block transactions:", err)
		return
	}

	// 過濾與該地址相關的交易
	for _, tx := range transactions {
		if tx.To == address || tx.From == address {
			p.storage.SaveTransaction(address, tx)
			p.notification.Notify(address, usecase.Transaction{
				BlockHash:   tx.BlockHash,
				BlockNumber: tx.BlockNumber,
				From:        tx.From,
				To:          tx.To,
				Value:       tx.Value,
			})
		}
	}
}

// PollForChanges 定期檢查區塊變化
func (p *EthereumParser) PollForChanges() {
	for {
		previousBlock := p.GetCurrentBlock()

		// 更新區塊，如果區塊有變化才繼續處理
		err := p.UpdateCurrentBlock()
		if err != nil {
			fmt.Println("Error updating current block:", err)
			continue
		}

		if p.currentBlock != previousBlock {
			fmt.Printf("New block detected: %d\n", p.currentBlock)
			// 檢查所有訂閱的地址並處理交易
			addresses := p.storage.GetSubscribedAddresses()
			for _, address := range addresses {
				p.FetchTransactionsForAddress(address)
			}
		}

		// 休眠 10 秒後再次檢查
		time.Sleep(10 * time.Second)
	}
}
