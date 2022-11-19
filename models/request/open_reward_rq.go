package request

type OpenRewardRq struct {
	GameId        string `form:"gameId" json:"gameId" binding:"required"`
	PlayerAddress string `form:"playerAddress" json:"playerAddress" binding:"required"`
}
