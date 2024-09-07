package usecase

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"parse_server/internal/domain/repository"
	"parse_server/internal/domain/usecase"
	"testing"

	repoMock "parse_server/internal/mock/repository"
	ucMock "parse_server/internal/mock/usecase"
)

func TestFetchBlockTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := repoMock.NewMockETHClient(ctrl)
	mockStorage := repoMock.NewMockStorage(ctrl)
	mockNotification := ucMock.NewMockNotification(ctrl)

	parser := NewEthereumParser(EthereumParserParam{
		Storage:      mockStorage,
		Notification: mockNotification,
		EthClient:    mockClient,
	})

	tests := []struct {
		name        string
		blockNumber string
		mockResult  json.RawMessage
		mockError   error
		expectedErr bool
		expectedTx  []repository.Transaction
	}{
		{
			name:        "Valid block with transactions",
			blockNumber: "0x10d4f",
			mockResult: json.RawMessage(`{
				"result": {
					"hash": "0xabc123",
					"number":"0x10d4f",
					"transactions": [
						{"from": "0xfrom1", "to": "0xto1", "value": "0x10"}
					]
				}
			}`),
			mockError:   nil,
			expectedErr: false,
			expectedTx: []repository.Transaction{
				{
					BlockHash:   "0xabc123",
					BlockNumber: "0x10d4f",
					From:        "0xfrom1",
					To:          "0xto1",
					Value:       "0x10",
				},
			},
		},
		{
			name:        "Error fetching block",
			blockNumber: "0x10d4f",
			mockResult:  nil,
			mockError:   errors.New("error calling Ethereum"),
			expectedErr: true,
			expectedTx:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 模擬 CallEthereum 函數的行為
			mockClient.EXPECT().CallEthereum("eth_getBlockByNumber", gomock.Any()).Return(tt.mockResult, tt.mockError)

			// 測試 fetchBlockTransactions
			transactions, err := parser.(*EthereumParser).fetchBlockTransactions(tt.blockNumber)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTx, transactions)
			}
		})
	}
}

func TestUpdateCurrentBlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := repoMock.NewMockETHClient(ctrl)
	mockStorage := repoMock.NewMockStorage(ctrl)
	mockNotification := ucMock.NewMockNotification(ctrl)

	parser := NewEthereumParser(EthereumParserParam{
		Storage:      mockStorage,
		Notification: mockNotification,
		EthClient:    mockClient,
	})

	tests := []struct {
		name          string
		mockResult    json.RawMessage
		mockError     error
		expectedErr   bool
		expectedBlock int
	}{
		{
			name: "Valid block number",
			mockResult: json.RawMessage(`{
				"result": "0x10d4f"
			}`),
			mockError:     nil,
			expectedErr:   false,
			expectedBlock: 68943, // 0x10d4f in decimal
		},
		{
			name:        "Error fetching block number",
			mockResult:  nil,
			mockError:   errors.New("error calling Ethereum"),
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 模擬 CallEthereum 行為
			mockClient.EXPECT().CallEthereum("eth_blockNumber", gomock.Any()).Return(tt.mockResult, tt.mockError)

			// 測試 UpdateCurrentBlock
			err := parser.(*EthereumParser).UpdateCurrentBlock()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBlock, parser.GetCurrentBlock())
			}
		})
	}
}

func TestFetchTransactionsForAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := repoMock.NewMockETHClient(ctrl)
	mockStorage := repoMock.NewMockStorage(ctrl)
	mockNotification := ucMock.NewMockNotification(ctrl)

	parser := NewEthereumParser(EthereumParserParam{
		Storage:      mockStorage,
		Notification: mockNotification,
		EthClient:    mockClient,
	})

	address := "0x123"

	mockClient.EXPECT().CallEthereum("eth_getBlockByNumber", gomock.Any()).Return(json.RawMessage(`{
		"result": {
			"hash": "0xabc123",
			"transactions": [
				{"from": "0x123", "to": "0x456", "value": "0x10"},
				{"from": "0x789", "to": "0x123", "value": "0x20"}
			]
		}
	}`), nil)

	// 模擬 SaveTransaction 和 Notify
	mockStorage.EXPECT().SaveTransaction(gomock.Any(), gomock.Any()).Times(2)
	mockNotification.EXPECT().Notify(gomock.Any(), gomock.Any()).Times(2)

	parser.(*EthereumParser).FetchTransactionsForAddress(address)
}

func TestSubscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := repoMock.NewMockETHClient(ctrl)
	mockStorage := repoMock.NewMockStorage(ctrl)
	mockNotification := ucMock.NewMockNotification(ctrl)

	parser := NewEthereumParser(EthereumParserParam{
		Storage:      mockStorage,
		Notification: mockNotification,
		EthClient:    mockClient,
	})

	address := "0x123"

	// 檢查 storage 是否調用了 SubscribeAddress
	mockStorage.EXPECT().SubscribeAddress(address).Times(1)

	// 調用 Subscribe 方法
	result := parser.Subscribe(address)

	// 檢查返回值
	assert.True(t, result)
}

func TestGetTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := repoMock.NewMockETHClient(ctrl)
	mockStorage := repoMock.NewMockStorage(ctrl)
	mockNotification := ucMock.NewMockNotification(ctrl)

	parser := NewEthereumParser(EthereumParserParam{
		Storage:      mockStorage,
		Notification: mockNotification,
		EthClient:    mockClient,
	})

	address := "0x123"

	// 模擬 Storage 返回的交易資料
	mockTransactions := []repository.Transaction{
		{
			BlockHash:   "0xabc123",
			BlockNumber: "100",
			From:        "0xfrom1",
			To:          "0xto1",
			Value:       "0x10",
		},
		{
			BlockHash:   "0xdef456",
			BlockNumber: "101",
			From:        "0xfrom2",
			To:          "0xto2",
			Value:       "0x20",
		},
	}

	// 模擬 GetTransactions 的返回結果
	mockStorage.EXPECT().GetTransactions(address).Return(mockTransactions).Times(1)

	// 調用 GetTransactions 方法
	result := parser.GetTransactions(address)

	// 檢查返回值
	expected := []usecase.Transaction{
		{
			BlockHash:   "0xabc123",
			BlockNumber: "100",
			From:        "0xfrom1",
			To:          "0xto1",
			Value:       "0x10",
		},
		{
			BlockHash:   "0xdef456",
			BlockNumber: "101",
			From:        "0xfrom2",
			To:          "0xto2",
			Value:       "0x20",
		},
	}

	assert.Equal(t, expected, result)
}
