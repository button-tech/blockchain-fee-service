package dto

type GetFeeRequest struct {
	FromAddress    string `json:"fromAddress"`
	Amount         string `json:"amount"`
	ReceiversCount int    `json:"receiversCount"`
}

type GetTokenFeeRequest struct {
	Address      string `json:"address"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
}
