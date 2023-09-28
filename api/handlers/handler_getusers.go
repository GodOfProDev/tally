package handlers

import (
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers()
	if err != nil {
		return responses.ErrGetting("users")
	}

	return responses.SuccessGotten(users)
}
