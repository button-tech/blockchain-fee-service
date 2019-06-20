package dto

type SharedApiResp struct {
	Fee                     int  `json:"fee"`
	Balance                 uint64  `json:"balance"`
	MaxAmountWithOptimalFee uint64  `json:"maxAmountWithOptimalFee"`
	MaxAmount               uint64  `json:"maxAmount,omitempty"`
	IsEnough                bool `json:"isEnough"`
	IsBadFee                bool `json:"isBadFee"`
	GasPrice                uint64  `json:"gasPrice,omitempty"`
	Gas                     uint64  `json:"gas,omitempty"`
}

type GetFeeResponse struct {
	*SharedApiResp
	Input                   int  `json:"input"`
	Output                  int  `json:"output"`
}

type GetEthFeeResponse struct {
	*SharedApiResp
}

type GetWavesAndStellarFeeResponse struct {
	*SharedApiResp
}

type GetTokenFeeResponse struct {
	*SharedApiResp
	TokenBalance            uint64  `json:"tokenBalance"`
}
