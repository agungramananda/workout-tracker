package exceptions

import "fmt"

type ClientError struct {
	Code    int
	Kind		string
	Message string
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("%s: Code=%d, Message=%s", e.Kind,e.Code, e.Message)
}

func NewClientError(code int, kind string, message string) *ClientError {
	if kind == "" {
		kind = "ClientError"
	}
	return &ClientError{
		Code:    code,
		Kind:    kind,
		Message: message,
	}
}