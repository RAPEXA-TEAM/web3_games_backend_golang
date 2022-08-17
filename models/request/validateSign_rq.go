package request

type ValidateSignRq struct {
	WalletAddress string `form:"walletAddress" json:"walletAddress" binding:"required"`
	Sign string `form:"sign" json:"sign" binding:"required"`
}
