package output

type GameRoundResultEvent struct {
	Winner string `form:"winner" json:"winner" binding:"required"`
}
