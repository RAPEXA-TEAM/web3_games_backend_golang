package response

type RockPaperActionRp struct {
	GameId        string `form:"gameId" json:"gameId" binding:"required"`
	PlayerAddress string `form:"playerAddress" json:"playerAddress" binding:"required"`
	Action        string `form:"action" json:"action" binding:"required"`
}
