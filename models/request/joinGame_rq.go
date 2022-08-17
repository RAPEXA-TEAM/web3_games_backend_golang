package request

type JoinGameRq struct {
	TxHash        string `form:"txHash" json:"txHash" binding:"required"`
	Value         string `form:"value" json:"value" binding:"required"`
	GameId        string `form:"gameId" json:"gameId" binding:"required"`
	PlayerAddress string `form:"playerAddress" json:"playerAddress" binding:"required"`
}
