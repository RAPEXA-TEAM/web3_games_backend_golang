package response

type ValidateSignRp struct {
	Success bool `form:"success" json:"success" binding:"required"`
}
