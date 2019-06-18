package dto

type GetFeeResponse struct {
	Fee                     int  `json:"fee"`
	Input                   int  `json:"input"`
	Output                  int  `json:"output"`
	Balance                 int  `json:"balance"`
	MaxAmount               int  `json:"maxAmount"`
	MaxAmountWithOptimalFee int  `json:"maxAmountWithOptimalFee"`
	IsEnough                bool `json:"isEnough"`
	IsBadFee                bool `json:"isBadFee"`
}
