package responses

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APIError) Error() string {
	return e.Msg
}
func ErrRequired(req string) APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    fmt.Sprintf("%v is required", req),
	}
}

func ErrInvalidId() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "invalid id",
	}
}

func ErrCreating(a string) APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    fmt.Sprintf("there was an issue creating the %v", a),
	}
}

func ErrNotFound(a string) APIError {
	return APIError{
		Status: fiber.StatusNotFound,
		Msg:    fmt.Sprintf("%v not found", a),
	}
}

func ErrUpdating(a string) APIError {
	return APIError{
		Status: fiber.StatusNotFound,
		Msg:    fmt.Sprintf("there was an issue updating the %v", a),
	}
}

func ErrParsingParams() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "there was an issue parsing params",
	}
}
