package handlers

import (
	"github.com/godofprodev/tally/api/models"
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleIncrement(c *fiber.Ctx) error {
	id, err := getId(c)
	if err != nil {
		return responses.ErrInvalidId()
	}

	guild, err := h.store.GetGuildById(id)
	if err != nil {
		return responses.ErrNotFound("guild")
	}

	params := new(models.IncrementGuildParams)

	if err := c.BodyParser(params); err != nil {
		return responses.ErrParsingParams()
	}

	if params.UserId == 0 {
		return responses.ErrRequired("userId")
	}

	user, err := h.store.GetUserById(params.UserId)
	if err != nil {
		return responses.ErrNotFound("user")
	}

	guild.Increment()
	user.Increment()

	err = h.store.UpdateGuild(guild)
	if err != nil {
		return responses.ErrUpdating("guild")
	}

	err = h.store.UpdateUser(user)
	if err != nil {
		return responses.ErrUpdating("user")
	}

	return responses.SuccessMessage("successfully incremented")
}
