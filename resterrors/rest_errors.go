package resterrors

import (
	"errors"
	"net/http"
)

// RestErr handle error
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"casuse"`
}

// NewError create new error
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError handle bad request error
func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError no user found
func NewNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError error due to internal server
func NewInternalServerError(msg string, err error) *RestErr {
	result := &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
