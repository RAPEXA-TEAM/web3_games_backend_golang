package request

type CreateGameRq struct {
	TxHash        string `form:"txHash" json:"txHash" binding:"required"`
	Value         string `form:"value" json:"value" binding:"required"`
	PlayerAddress string `form:"playerAddress" json:"playerAddress" binding:"required"`
}
