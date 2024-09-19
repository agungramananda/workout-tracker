package exceptions

import (
	"fmt"
	"net/http"
)

type InvalidRequestError struct {
	ClientError *ClientError
	Cause       error
}

func (e *InvalidRequestError) Error() string {
	return e.ClientError.Error() + fmt.Sprintf(", Cause=%v", e.Cause)
}

func NewInvalidRequestError(message string, cause error) *InvalidRequestError {
	return &InvalidRequestError{
		ClientError: NewClientError(http.StatusBadRequest, "InvalidRequestError", message),
		Cause:       cause,
	}
}