package responses

type UtxoResponse struct {
	Utxo []Utxo `json:"utxo"`
}

type Utxo struct {
	Address       string `json:"address"`
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	ScriptPubKey  string `json:"scriptPubKey"`
	Amount        string `json:"amount"`
	Satoshis      int    `json:"satoshis"`
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
}
