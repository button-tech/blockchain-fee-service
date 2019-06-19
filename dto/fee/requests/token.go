package requests

type TokenGasLimitRequest struct {
	ToAddress    string `json:"toAddress"`
	TokenAddress string `json:"tokenAddress"`
	Amount       string `json:"Amount"`
}
