package responses

import "time"

type BitcoinFeeResponse struct {
	FastestFee  int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
}

type LitecoinFeeResponse struct {
	Name             string    `json:"name"`
	Height           int       `json:"height"`
	Hash             string    `json:"hash"`
	Time             time.Time `json:"time"`
	LatestURL        string    `json:"latest_url"`
	PreviousHash     string    `json:"previous_hash"`
	PreviousURL      string    `json:"previous_url"`
	PeerCount        int       `json:"peer_count"`
	UnconfirmedCount int       `json:"unconfirmed_count"`
	HighFeePerKb     int       `json:"high_fee_per_kb"`
	MediumFeePerKb   int       `json:"medium_fee_per_kb"`
	LowFeePerKb      int       `json:"low_fee_per_kb"`
	LastForkHeight   int       `json:"last_fork_height"`
	LastForkHash     string    `json:"last_fork_hash"`
}

type EthereumFeeResponse struct {
	GasPrice int `json:"gasPrice"`
}
