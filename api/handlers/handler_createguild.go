package handlers

import (
	"github.com/godofprodev/tally/api/models"
	"github.com/godofprodev/tally/api/responses"
	"github.com/gofiber/fiber/v2"
)

func (h Handlers) HandleCreateGuild(c *fiber.Ctx) error {
	params := new(models.CreateGuildParams)

	if err := c.BodyParser(params); err != nil {
		return responses.ErrParsingParams()
	}

	if params.ServerId == 0 {
		return responses.ErrRequired("serverId")
	}
	if params.ChannelId == 0 {
		return responses.ErrRequired("channelId")
	}

	guild := models.NewGuild(params)

	err := h.store.CreateGuild(guild)
	if err != nil {
		return responses.ErrCreatingUser()
	}

	return responses.SuccessCreated(guild)
}
