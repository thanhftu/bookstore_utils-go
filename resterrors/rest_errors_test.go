package resterrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewError("This is error for test")
	assert.NotNil(t, err)
	assert.EqualValues(t, "This is error for test", err.Error())
}
func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "internal server error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, "database error", err.Causes[0])
	assert.EqualValues(t, 1, len(err.Causes))

}
