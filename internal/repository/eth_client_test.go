package repository

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"parse_server/internal/domain"
	"testing"
)

func TestCallEthereum(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name           string
		method         string
		params         []interface{}
		mockResponse   string
		mockStatus     int
		mockError      error
		expectedResult []byte
		expectedErr    bool
	}{
		{
			name:           "Successful request",
			method:         "eth_blockNumber",
			params:         []interface{}{},
			mockResponse:   `{"jsonrpc":"2.0","id":1,"result":"0x10d4f"}`,
			mockStatus:     200,
			expectedResult: []byte(`{"jsonrpc":"2.0","id":1,"result":"0x10d4f"}`),
			expectedErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 模擬 HTTP 請求
			httpmock.RegisterResponder("POST", domain.DefaultURL,
				func(req *http.Request) (*http.Response, error) {
					if tt.mockError != nil {
						return nil, tt.mockError
					}
					resp := httpmock.NewStringResponse(tt.mockStatus, tt.mockResponse)
					return resp, nil
				})

			// 創建 Client 並調用 CallEthereum
			client := Client{}
			result, err := client.CallEthereum(tt.method, tt.params)

			// 驗證結果
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
