package resterrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewRestError("test error", http.StatusNotImplemented, "not_implement")
	assert.NotNil(t, err)
	assert.EqualValues(t, "message: test error - status: 501 - error: not_implement - causes: []", err.Error())
}

func TestNewRestError(t *testing.T) {
	err1 := NewRestError("error 1 without cause", http.StatusNotImplemented, "error_1")
	assert.NotNil(t, err1)
	assert.EqualValues(t, http.StatusNotImplemented, err1.Status)
	assert.EqualValues(t, "error 1 without cause", err1.Message)
	assert.EqualValues(t, "error_1", err1.Error)
	assert.Nil(t, err1.Causes)

	err2 := NewRestError("error 2 with cause", http.StatusNotImplemented, "error_2", errors.New("causing error"))
	assert.NotNil(t, err2)
	assert.EqualValues(t, http.StatusNotImplemented, err2.Status)
	assert.EqualValues(t, "error 2 with cause", err2.Message)
	assert.EqualValues(t, "error_2", err2.Error)
	assert.NotNil(t, err2.Causes)
	assert.EqualValues(t, 1, len(err2.Causes()))
}
func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "internal server error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, "database error", err.Causes()[0])
	assert.EqualValues(t, 1, len(err.Causes()))

}
