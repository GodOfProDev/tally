package handlers

import (
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleReset(c *fiber.Ctx) error {
	id, err := getId(c)
	if err != nil {
		return responses.ErrInvalidId()
	}

	guild, err := h.store.GetGuildById(id)
	if err != nil {
		return responses.ErrNotFound("guild")
	}

	guild.CurrentCount = 0

	err = h.store.UpdateGuild(guild)
	if err != nil {
		return responses.ErrUpdating("guild")
	}

	return nil
}
