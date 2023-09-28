package handlers

import (
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleGetGuilds(c *fiber.Ctx) error {
	guilds, err := h.store.GetGuilds()
	if err != nil {
		return responses.ErrNotFound("guilds")
	}

	return responses.SuccessGotten(guilds)
}
