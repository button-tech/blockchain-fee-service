package handlers

import (
	"dev.azure.com/fee-service/dto/errors"
	"dev.azure.com/fee-service/dto/fee/responses"
	"log"
	"net/http"
)

func handleError(e interface{}) (bool, int, errors.ApiError) {
	switch resp := e.(type) {
	case responses.ResponseError:
		if resp.ApiError != nil {
			log.Println(resp.ApiError)
			return true, http.StatusNotImplemented, errors.ApiError{
				ExceptionId: http.StatusNotImplemented,
				Error:       resp.ApiError,
			}
		}
		if resp.Error != nil {
			log.Println(resp.Error)
			return true, http.StatusInternalServerError, errors.ApiError{
				ExceptionId: http.StatusInternalServerError,
				Error:       "Internal Server Error",
			}
		}
		return false, 0, errors.ApiError{}
	case errors.BadRequest:
		if resp.Error != nil {
			log.Println(resp.Error)
			return true, http.StatusBadRequest, errors.ApiError{
				ExceptionId: http.StatusBadRequest,
				Error:       "Bad Request",
			}
		}
		return false, 0, errors.ApiError{}
	case error:
		log.Println(resp)
		return true, http.StatusInternalServerError, errors.ApiError{
			ExceptionId: http.StatusInternalServerError,
			Error:       "Internal Server Error",
		}
	default:
		return false, 0, errors.ApiError{}
	}
}
