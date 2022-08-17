package output

type GameEndResultEvent struct {
	Winner string `form:"winner" json:"winner" binding:"required"`
}
