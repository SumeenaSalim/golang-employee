package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type Error struct {
	Message string       `json:"message,omitempty" binding:"required"`
	Details ErrorDetails `json:"details,omitempty"`
}

type ErrorData interface {
	GetErrorDetails() []ErrorDetail
}

type ErrorDetail struct {
	Field string `json:"field,omitempty"`
	Error string `json:"error,omitempty"`
}

func (err ErrorDetail) GetErrorDetails() []ErrorDetail {
	return []ErrorDetail{err}
}

type ErrorDetails []ErrorDetail

func (errs ErrorDetails) GetErrorDetails() []ErrorDetail {
	return errs
}

func JSONSuccess(c *gin.Context, httpStatus int, data interface{}, optionalMessage ...string) {
	fmt.Println("Sending Success Response...")
	status := "success"
	res := JSONResponse{
		Status: status,
		Data:   data,
	}

	if len(optionalMessage) > 0 {
		res.Message = optionalMessage[0]
	}

	c.JSON(httpStatus, res)
}

func JSONError(c *gin.Context, httpStatus int, message string, optionalError ...ErrorData) {
	status := "error"
	if message == "" && optionalError != nil {
		message = optionalError[0].GetErrorDetails()[0].Error
	}

	res := JSONResponse{
		Status: status,
		Error: &Error{
			Message: message,
		},
	}

	if len(optionalError) > 0 {
		res.Error.Details = optionalError[0].GetErrorDetails()
	}

	c.JSON(httpStatus, res)
}
