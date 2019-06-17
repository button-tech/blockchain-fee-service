package fee

import (
	"dev.azure.com/fee-service/dto/fee/responses"
	"encoding/json"
	"github.com/imroc/req"
	"log"
	"net/http"
)

type apiResponse struct {
	Result     interface{} `json:"result"`
	StatusCode int         `json:"statusCode"`
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
				Result:     err,
				StatusCode: http.StatusInternalServerError,
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

	statusCode := response.Response().StatusCode
	if respErr != nil {
		return apiResponse{
			Result:     respErr,
			StatusCode: statusCode,
		}
	}

	return apiResponse{
		Result:     response,
		StatusCode: statusCode,
	}
}

func (res *apiResponse) response(v interface{}) responses.Response {
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		if err := res.Result.(*req.Resp).ToJSON(&v); err != nil {
			log.Println(err.Error())
			return responses.Response{
				Error: err.Error(),
			}
		}
		return responses.Response{
			Result: v,
		}
	} else {
		return responses.Response{
			Error: v,
		}
	}

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
