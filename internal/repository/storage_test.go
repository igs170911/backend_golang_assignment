package repository

import (
	"github.com/stretchr/testify/assert"
	domainRepo "parse_server/internal/domain/repository"
	"testing"
)

func TestMemoryStorage_SaveAndGetTransactions(t *testing.T) {
	tests := []struct {
		name          string
		address       string
		transactions  []domainRepo.Transaction
		expectedCount int
	}{
		{
			name:    "Single transaction",
			address: "0x123",
			transactions: []domainRepo.Transaction{
				{
					BlockHash:   "0xhash1",
					BlockNumber: "100",
					From:        "0xfrom1",
					To:          "0xto1",
					Value:       "0x10",
				},
			},
			expectedCount: 1,
		},
		{
			name:    "Multiple transactions",
			address: "0x456",
			transactions: []domainRepo.Transaction{
				{
					BlockHash:   "0xhash2",
					BlockNumber: "101",
					From:        "0xfrom2",
					To:          "0xto2",
					Value:       "0x20",
				},
				{
					BlockHash:   "0xhash3",
					BlockNumber: "102",
					From:        "0xfrom3",
					To:          "0xto3",
					Value:       "0x30",
				},
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初始化內存存儲
			storage := NewMemoryStorage()

			// 保存交易
			for _, tx := range tt.transactions {
				storage.SaveTransaction(tt.address, tx)
			}

			// 獲取交易並檢查結果
			result := storage.GetTransactions(tt.address)
			assert.Equal(t, tt.expectedCount, len(result))
			assert.Equal(t, tt.transactions, result)
		})
	}
}

func TestMemoryStorage_SubscribeAddress(t *testing.T) {
	tests := []struct {
		name           string
		addresses      []string
		expectedResult []string
	}{
		{
			name:           "Single address subscription",
			addresses:      []string{"0x123"},
			expectedResult: []string{"0x123"},
		},
		{
			name:           "Multiple address subscription",
			addresses:      []string{"0x123", "0x456", "0x789"},
			expectedResult: []string{"0x123", "0x456", "0x789"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初始化內存存儲
			storage := NewMemoryStorage()

			// 訂閱地址
			for _, addr := range tt.addresses {
				storage.SubscribeAddress(addr)
			}

			// 獲取訂閱的地址
			result := storage.GetSubscribedAddresses()

			// 檢查結果
			assert.ElementsMatch(t, tt.expectedResult, result)
		})
	}
}

func TestMemoryStorage_GetTransactionsForNonExistentAddress(t *testing.T) {
	storage := NewMemoryStorage()

	// 嘗試獲取一個不存在地址的交易
	result := storage.GetTransactions("0xnonexistent")

	// 應該返回空的交易列表
	assert.Empty(t, result)
}
