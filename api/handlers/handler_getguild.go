package handlers

import (
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetGuild(c *fiber.Ctx) error {
	id, err := getId(c)
	if err != nil {
		return responses.ErrInvalidId()
	}

	guild, err := h.store.GetGuildById(id)
	if err != nil {
		return responses.ErrNotFound("guild")
	}

	return responses.SuccessGotten(guild)
}
