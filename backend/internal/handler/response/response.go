package response

import "github.com/gin-gonic/gin"

type APIResponse[T any] struct {
	Success bool      `json:"success"`
	Data    T         `json:"data,omitempty"`
	Error   *APIError `json:"error,omitempty"`
	Meta    *Meta     `json:"meta,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	RequestID string `json:"request_id,omitempty"`
}

func NewError(code, message string) APIError {
	return APIError{
		Code:    code,
		Message: message,
	}
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, APIResponse[any]{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, err APIError) {
	c.JSON(status, APIResponse[any]{
		Success: false,
		Error:   &err,
	})
}
