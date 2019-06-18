package dto

type GetFeeRequest struct {
	FromAddress    string  `json:"fromAddress"`
	Amount         float64 `json:"amount"`
	ReceiversCount int     `json:"receiversCount"`
}
