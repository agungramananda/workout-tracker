package exceptions

import "net/http"

type AuthorizationError struct {
	ClientError *ClientError
}

func (e *AuthorizationError) Error() string {
	return e.ClientError.Error()
}

func NewAuthorizationError(message string) *AuthorizationError {
	return &AuthorizationError{
		ClientError: NewClientError(http.StatusForbidden, "AuthorizationError", message),
	}
}