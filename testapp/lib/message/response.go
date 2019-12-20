package message

import "github.com/joshia/automated-api-test-service/testapp/config/apperror"

type ErrorResponse struct {
	Code	string	`json:"code"`
	Message	string	`json:"message"`
}

func SetErrorResponse(err *apperror.V1Error) *ErrorResponse {
	errRes := &ErrorResponse{
		Code:    err.ErrorCode,
		Message: err.PublicMessage,
	}
	return errRes
}
