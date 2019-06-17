package responses

type BitcoinFeeResponse struct {
	FastestFee  int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
}
