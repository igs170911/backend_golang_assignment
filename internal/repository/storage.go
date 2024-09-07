package repository

import (
	"parse_server/internal/domain/repository"
)

// MemoryStorage 實現了 Storage interface
type MemoryStorage struct {
	addresses    map[string]bool
	transactions map[string][]repository.Transaction
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		addresses:    make(map[string]bool),
		transactions: make(map[string][]repository.Transaction),
	}
}

func (m *MemoryStorage) SaveTransaction(address string, tx repository.Transaction) {
	m.transactions[address] = append(m.transactions[address], tx)
}

func (m *MemoryStorage) GetTransactions(address string) []repository.Transaction {
	return m.transactions[address]
}

func (m *MemoryStorage) GetSubscribedAddresses() []string {
	var addresses []string
	for addr := range m.addresses {
		addresses = append(addresses, addr)
	}
	return addresses
}

func (m *MemoryStorage) SubscribeAddress(address string) {
	m.addresses[address] = true
}
