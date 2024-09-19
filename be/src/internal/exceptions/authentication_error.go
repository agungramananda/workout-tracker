package exceptions

import "net/http"

type AuthenticationError struct {
	ClientError *ClientError
}

func (e *AuthenticationError) Error() string {
	return e.ClientError.Error()
}

func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{
		ClientError: NewClientError(http.StatusUnauthorized, "AuthenticationError", message),
	}
}