package dto

type SharedApiReq struct {
	FromAddress    string `json:"fromAddress"`
	Amount         string `json:"amount"`
}

type GetFeeRequest struct {
	*SharedApiReq
	ReceiversCount int    `json:"receiversCount"`
}

type GetTokenFeeRequest struct {
	*SharedApiReq
	TokenAddress string `json:"tokenAddress"`
}
