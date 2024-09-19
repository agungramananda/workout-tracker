package utils

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

func NewResponse(status int, message string, data []interface{}) *Response {
	if data == nil {
		return &Response{
			Status:  status,
			Message: message,
		}
	}
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewErrorResponse(status int, error string, message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Error:   error,
		Message: message,
	}
}