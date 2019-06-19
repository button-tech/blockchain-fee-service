package responses

type CurrencyBalanceResponse struct {
	Balance string `json:"balance"`
}

type WavesBalanceResponse struct {
	Address       string `json:"address"`
	Confirmations int    `json:"confirmations"`
	Balance       int    `json:"balance"`
}
