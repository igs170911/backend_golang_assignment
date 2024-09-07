package payload

type SubscribeReq struct {
	Address string `json:"address" binding:"required"`
}
