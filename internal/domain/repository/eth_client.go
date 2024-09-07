package repository

// EthereumRPCResponse JSON-RPC 與 Ethereum 節點通訊
type EthereumRPCResponse struct {
	ID      int    `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}

// Block 定義 JSON RPC 區塊返回的結構

type BlockResult struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Block  `json:"result"`
}

type Block struct {
	Difficulty       string            `json:"difficulty"`
	ExtraData        string            `json:"extraData"`
	GasLimit         string            `json:"gasLimit"`
	GasUsed          string            `json:"gasUsed"`
	Hash             string            `json:"hash"`
	LogsBloom        string            `json:"logsBloom"`
	Miner            string            `json:"miner"`
	MixHash          string            `json:"mixHash"`
	Nonce            string            `json:"nonce"`
	Number           string            `json:"number"`
	ParentHash       string            `json:"parentHash"`
	ReceiptsRoot     string            `json:"receiptsRoot"`
	Sha3Uncles       string            `json:"sha3Uncles"`
	Size             string            `json:"size"`
	StateRoot        string            `json:"stateRoot"`
	Timestamp        string            `json:"timestamp"`
	TotalDifficulty  string            `json:"totalDifficulty"`
	Transactions     []TransactionItem `json:"transactions"`
	TransactionsRoot string            `json:"transactionsRoot"`
	Uncles           []string          `json:"uncles"`
}

type TransactionItem struct {
	BlockHash            *string `json:"blockHash"`            // 區塊的哈希值，當交易在等待中時為null
	BlockNumber          *string `json:"blockNumber"`          // 區塊編號，當交易在等待中時為null
	From                 string  `json:"from"`                 // 發送者的地址
	Gas                  string  `json:"gas"`                  // 發送者提供的gas（十六進位編碼）
	GasPrice             string  `json:"gasPrice"`             // 發送者提供的gas價格（以wei為單位，十六進位編碼）
	MaxFeePerGas         *string `json:"maxFeePerGas"`         // 設定的每單位 gas 的最大費用（可選，可能為nil）
	MaxPriorityFeePerGas *string `json:"maxPriorityFeePerGas"` // 設定的優先級 gas 費用的最大值（可選，可能為nil）
	Hash                 string  `json:"hash"`                 // 交易的哈希值
	Input                string  `json:"input"`                // 與交易一起發送的數據
	Nonce                string  `json:"nonce"`                // 發送者在此交易之前發送的交易數（十六進位編碼）
	To                   *string `json:"to"`                   // 接收者的地址，當是合約創建交易時為null
	TransactionIndex     *string `json:"transactionIndex"`     // 交易索引位置，當是等待中的交易時為null
	Value                string  `json:"value"`                // 轉移的金額（以wei為單位，十六進位編碼）
	Type                 string  `json:"type"`                 // 交易的類型
	AccessList           []any   `json:"accessList"`           // 計劃訪問的地址和存儲鍵列表
	ChainId              *string `json:"chainId"`              // 交易的鏈ID（若有）
	V                    string  `json:"v"`                    // 簽名中的標準化V字段
	R                    string  `json:"r"`                    // 簽名中的R字段
	S                    string  `json:"s"`                    // 簽名中的S字段
}

type ETHClient interface {
	CallEthereum(method string, params []any) ([]byte, error)
}
