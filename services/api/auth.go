package api

import (
	"dev.azure.com/fee-service/dto/fee/responses"
	"encoding/json"
	"github.com/imroc/req"
	"log"
	"net/http"
)

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
	if method == "GET" {
		response, respErr = req.Get(uri, authHeader)
	} else if method == "POST" {
		response, respErr = req.Post(uri, authHeader, data)
	} else if method == "PATCH" {
		response, respErr = req.Patch(uri, authHeader, data)
	} else if method == "DELETE" {
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
		//var e interface{}
		//if err := res.ApiError.(*req.Resp).ToJSON(&e); err != nil {
		//	return responses.ResponseError{
		//		Error: err,
		//	}
		//}
		return responses.ResponseError{
			ApiError: res.ApiError.(*req.Resp).String(),
		}
	}
	if res.Error != nil {
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
		break
	default:
		result, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return result, nil
		break
	}
	return nil, nil
}
