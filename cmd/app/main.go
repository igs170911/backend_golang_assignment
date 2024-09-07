package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parse_server/internal/delivery/http/payload"
	"parse_server/internal/delivery/http/request"
	domainUC "parse_server/internal/domain/usecase"
	"parse_server/internal/repository"
	"parse_server/internal/usecase"
)

var P domainUC.Parser

func main() {
	// 初始化 Storage 和 Notification
	storage := repository.NewMemoryStorage()
	notification := usecase.MustNotification()
	ethClient := repository.MustETHClient(repository.ClientParam{})

	// 初始化 Parser
	P = usecase.NewEthereumParser(usecase.EthereumParserParam{
		Storage:      storage,
		Notification: notification,
		EthClient:    ethClient,
	})

	// 開始檢查區塊變化
	go P.PollForChanges()

	// 使用 gin.New() 創建 Gin 引擎
	r := gin.New()

	// 設定路由，只允許 POST 請求
	r.POST("/subscribe", SubscribeHandler)

	// 啟動伺服器
	r.Run(":8080") // 預設監聽在 8080 埠

}

// SubscribeHandler 處理訂閱請求
func SubscribeHandler(c *gin.Context) {
	// 定義 request 結構
	var req payload.SubscribeReq

	// 解析 JSON 請求體並驗證
	if err := request.ShouldBindJSON(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 執行訂閱操作
	status := P.Subscribe(req.Address)
	if !status {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to subscribe"})
		return
	}

	// 返還成功訊息
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
