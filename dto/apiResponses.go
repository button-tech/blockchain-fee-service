package dto

type SharedEthBasedResp struct {
	GasPrice uint64 `json:"gasPrice"`
	Gas      uint64 `json:"gas"`
}

type SharedApiResp struct {
	Fee                     int    `json:"fee"`
	Balance                 uint64 `json:"balance"`
	MaxAmountWithOptimalFee uint64 `json:"maxAmountWithOptimalFee"`
	MaxAmount               uint64 `json:"maxAmount"`
	IsEnough                bool   `json:"isEnough"`
	IsBadFee                bool   `json:"isBadFee"`
}

type GetFeeResponse struct {
	*SharedApiResp
	Input  int `json:"input"`
	Output int `json:"output"`
}

type GetEthFeeResponse struct {
	*SharedApiResp
	*SharedEthBasedResp
}

type GetWavesAndStellarFeeResponse struct {
	*SharedApiResp
}

type GetTokenFeeResponse struct {
	*SharedApiResp
	*SharedEthBasedResp
	TokenBalance uint64 `json:"tokenBalance"`
}
