package api

import (
	"encoding/json"
	"github.com/button-tech/blockchain-fee-service/dto/fee/responses"
	"github.com/imroc/req"
	"log"
	"net/http"
	"os"
)

var nodeUrl = os.Getenv("NODE_URL")

type apiResponse struct {
	Result   interface{} `json:"result"`
	ApiError interface{} `json:"apiError"`
	Error    error       `json:"error"`
}

func apiCall(method, apiUrl, path string, data interface{}) apiResponse {
	var response *req.Resp
	var respErr error
	authHeader := req.Header{
		"Content-Type": "application/json",
	}

	uri := apiUrl + path
	if data != nil {
		var err error
		data, err = serialize(data)
		if err != nil {
			return apiResponse{
				Error: err,
			}
		}
	}

	switch method {
	case "GET":
		response, respErr = req.Get(uri, authHeader)
	case "POST":
		response, respErr = req.Post(uri, authHeader, data)
	case "PUT":
		response, respErr = req.Put(uri, authHeader, data)
	case "PATCH":
		response, respErr = req.Patch(uri, authHeader, data)
	case "DELETE":
		response, respErr = req.Delete(uri, authHeader)
	}

	if respErr != nil {
		return apiResponse{
			Error: respErr,
		}
	}

	if statusCode := response.Response().StatusCode; statusCode == http.StatusOK || statusCode == http.StatusCreated {
		return apiResponse{
			Result: response,
		}
	}

	return apiResponse{
		ApiError: response,
	}
}

func (res *apiResponse) response(v interface{}) responses.ResponseError {
	if res.ApiError != nil {
		return responses.ResponseError{
			ApiError: res.ApiError.(*req.Resp).String(),
		}
	} else if res.Error != nil {
		return responses.ResponseError{
			Error: res.Error,
		}
	}

	if err := res.Result.(*req.Resp).ToJSON(&v); err != nil {
		log.Println(err.Error())
		return responses.ResponseError{
			Error: err,
		}
	}
	return responses.ResponseError{}
}

func serialize(data interface{}) ([]byte, error) {
	switch data.(type) {
	case []byte:
		return data.([]byte), nil
	default:
		result, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}
