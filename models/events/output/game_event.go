package output

type GameEvent struct {
	GameId          string `form:"gameId" json:"gameId" binding:"required"`
	PlayerAddresses string `form:"playerAddresses" json:"playerAddresses" binding:"required"`
	Status          int    `form:"status" json:"status" binding:"required"`
}
