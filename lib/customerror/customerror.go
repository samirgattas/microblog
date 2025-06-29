package customerror

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Message: %s, Status: %d", e.Message, e.StatusCode)
}

func NewNotFoundError(entity string) error {
	return NotFoundError{
		CustomError: CustomError{
			Message:    fmt.Sprintf("%s not found", entity),
			StatusCode: http.StatusNotFound,
		},
	}
}

func NewBadRequestError(msg string) error {
	return BadRequestError{
		CustomError{
			Message:    msg,
			StatusCode: http.StatusBadRequest,
		},
	}
}

func NewInternalServerError(msg string) error {
	return InternalServerError{
		CustomError{
			Message:    msg,
			StatusCode: http.StatusBadRequest,
		},
	}
}

type NotFoundError struct {
	CustomError
}

type BadRequestError struct {
	CustomError
}

type InternalServerError struct {
	CustomError
}
