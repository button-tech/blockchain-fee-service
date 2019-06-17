package handlers

import (
	"dev.azure.com/moon-pay/dto/errors"
	"dev.azure.com/moon-pay/dto/moonpay/responses"
	"log"
	"net/http"
)

func handleError(resp interface{}, bindStruct interface{}) (bool, int, errors.ApiError) {
	switch resp := resp.(type) {
	case responses.Response:
		if resp.Result == nil {
			log.Println(resp.Error)
			return true, http.StatusNotImplemented, errors.ApiError{
				ExceptionId: errors.ExceptionIdMoonPayError,
				Message:     resp.Error.Message,
			}
		}
		if bindStruct == nil {
			return false, 0, errors.ApiError{}
		}
		if err := resp.ToStruct(&bindStruct); err != nil {
			log.Println(err)
			return true, http.StatusNotImplemented, errors.ApiError{
				ExceptionId: errors.ExceptionIdMoonPayError,
				Message:     err.Error(),
			}
		}
		return false, 0, errors.ApiError{}
	case errors.BadRequest:
		if resp.Error != nil {
			log.Println(resp.Error)
			return true, http.StatusBadRequest, errors.ApiError{
				ExceptionId: errors.ExceptionIdBadRequest,
				Message:     resp.Message,
			}
		}
		return false, 0, errors.ApiError{}
	case errors.NotFound:
		if resp.Error != nil {
			log.Println(resp.Error)
			return true, http.StatusNotFound, errors.ApiError{
				ExceptionId: errors.ExceptionIdNotExists,
				Message:     resp.Message,
			}
		}
		return false, 0, errors.ApiError{}
	case errors.Conflict:
		log.Println(resp.Error)
		if resp.Error != nil {
			return true, http.StatusConflict, errors.ApiError{
				ExceptionId: errors.ExceptionIdExists,
				Message:     resp.Message,
			}
		}
		return false, 0, errors.ApiError{}
	case error:
		log.Println(resp)
		return true, http.StatusInternalServerError, errors.ApiError{
			ExceptionId: http.StatusInternalServerError,
			Message:     "Internal Server Error",
		}
	default:
		return false, 0, errors.ApiError{}
	}
}
