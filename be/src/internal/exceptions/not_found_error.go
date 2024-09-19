package exceptions

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ClientError *ClientError
	Cause       error
}

func (e *NotFoundError) Error() string {
	return e.ClientError.Error() + fmt.Sprintf(", Cause=%v", e.Cause)
}

func NewNotFoundError(message string, cause error) *NotFoundError {
	return &NotFoundError{
		ClientError: NewClientError(http.StatusNotFound, "NotFoundError", message),
		Cause:       cause,
	}
}