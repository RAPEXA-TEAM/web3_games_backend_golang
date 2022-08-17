package response

type RequestNonceRp struct {
	Nonce string `form:"nonce" json:"nonce" binding:"required"`
}
