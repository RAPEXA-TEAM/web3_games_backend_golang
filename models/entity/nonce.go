package entity

type Nonce struct {
	Id            int    `json:"id"`
	WalletAddress string `json:"wallet_address"`
	Nonce         string `json:"nonce"`
}
