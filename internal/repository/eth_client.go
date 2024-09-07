package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"parse_server/internal/domain"
	"parse_server/internal/domain/repository"
	"strings"
)

type ClientParam struct{}

type Client struct{}

func (c Client) CallEthereum(method string, params []interface{}) ([]byte, error) {
	rpcBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
	jsonBody, err := json.Marshal(rpcBody)
	if err != nil {
		return []byte{}, err
	}

	resp, err := http.Post(domain.DefaultURL, "application/json", strings.NewReader(string(jsonBody)))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// 讀取 HTTP response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func MustETHClient(_ ClientParam) repository.ETHClient {
	return &Client{}
}
