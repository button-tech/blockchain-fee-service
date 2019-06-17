package responses

type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"response"`
}
