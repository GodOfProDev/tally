package handlers

import (
	"github.com/godofprodev/tally/api/models"
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleCreateUser(c *fiber.Ctx) error {
	params := new(models.CreateUserParams)

	if err := c.BodyParser(params); err != nil {
		return responses.ErrParsingParams()
	}

	if params.UserId == 0 {
		return responses.ErrRequired("userId")
	}

	user := models.NewUser(params)

	err := h.store.CreateUser(user)
	if err != nil {
		return responses.ErrCreating("user")
	}

	return responses.SuccessCreated(user)
}
