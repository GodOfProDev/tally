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

func ErrCreatingUser() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue creating the user",
	}
}

func ErrGettingUser() APIError {
	return APIError{
		Status: fiber.StatusNotFound,
		Msg:    "there was an issue getting the user",
	}
}

func ErrGettingUsers() APIError {
	return APIError{
		Status: fiber.StatusNotFound,
		Msg:    "there was an issue getting the users",
	}
}

func ErrParsingParams() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "there was an issue parsing params",
	}
}
