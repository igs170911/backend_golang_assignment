.PHONY: mock-gen
mock-gen: # 建立 mock 資料
	mockgen -source=./internal/domain/repository/eth_client.go -destination=./internal/mock/repository/eth_client.go -package=mock
	mockgen -source=./internal/domain/repository/storage.go -destination=./internal/mock/repository/storage.go -package=mock
	mockgen -source=./internal/domain/usecase/notification.go -destination=./internal/mock/usecase/notification.go -package=mock
	mockgen -source=./internal/domain/usecase/parse.go -destination=./internal/mock/usecase/parse.go -package=mock

