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

// NewRestError Create new RestErr
func NewRestError(msg string, status int, err string, causes ...error) *RestErr {
	RestErr := &RestErr{
		Message: msg,
		Status:  status,
		Error:   err,
	}
	if len(causes) != 0 {
		for _, err := range causes {

			RestErr.Causes = append(RestErr.Causes, err.Error())
		}
	}
	return RestErr
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
