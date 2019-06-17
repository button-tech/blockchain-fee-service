package errors

import e "errors"

type ApiError struct {
	ExceptionId int    `json:"exceptionId"`
	Message     string `json:"message"`
}

type BadRequest struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type Conflict struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type NotFound struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func CustomError(message string) error {
	return e.New(message)
}
