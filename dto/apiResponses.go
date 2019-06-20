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

type GetEthFeeResponse struct {
	Fee                     int  `json:"fee"`
	GasPrice                int  `json:"gasPrice"`
	Gas                     int  `json:"gas"`
	Balance                 int  `json:"balance"`
	MaxAmount               int  `json:"maxAmount"`
	MaxAmountWithOptimalFee int  `json:"maxAmountWithOptimalFee"`
	IsEnough                bool `json:"isEnough"`
	IsBadFee                bool `json:"isBadFee"`
}

type GetWavesAndStellarFeeResponse struct {
	Fee                     int  `json:"fee"`
	Balance                 int  `json:"balance"`
	MaxAmountWithOptimalFee int  `json:"maxAmountWithOptimalFee"`
	IsEnough                bool `json:"isEnough"`
	IsBadFee                bool `json:"isBadFee"`
}

type GetTokenFeeResponse struct {
	Fee                     int  `json:"fee"`
	GasPrice                int  `json:"gasPrice"`
	Gas                     int  `json:"gas"`
	Balance                 int  `json:"balance"`
	TokenBalance            int  `json:"tokenBalance"`
	MaxAmountWithOptimalFee int  `json:"maxAmountWithOptimalFee"`
	IsEnough                bool `json:"isEnough"`
	IsBadFee                bool `json:"isBadFee"`
}
