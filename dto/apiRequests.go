package dto

type GetFeeRequest struct {
	FromAddress    string `json:"fromAddress"`
	Amount         string `json:"amount"`
	ReceiversCount int    `json:"receiversCount"`
}
