package responses

type ResponseError struct {
	ApiError interface{} `json:"apiError"`
	Error    interface{} `json:"error"`
}
