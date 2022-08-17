package request

type RequestNonceRq struct {
	WalletAddress string `form:"walletAddress" json:"walletAddress" binding:"required"`
}
