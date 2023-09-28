package handlers

import (
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetUser(c *fiber.Ctx) error {
	id, err := getId(c)
	if err != nil {
		return responses.ErrInvalidId()
	}

	user, err := h.store.GetUserById(id)
	if err != nil {
		return responses.ErrGetting("user")
	}

	return responses.SuccessGotten(user)
}
