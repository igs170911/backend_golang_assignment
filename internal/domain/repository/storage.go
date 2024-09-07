package repository

// Storage interface
type Storage interface {
	SaveTransaction(address string, tx Transaction)
	GetTransactions(address string) []Transaction
	GetSubscribedAddresses() []string
	SubscribeAddress(address string)
}

type Transaction struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
}
