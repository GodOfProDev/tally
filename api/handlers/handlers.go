package handlers

import (
	"github.com/godofprodev/tally/api/storage"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handlers struct {
	store storage.Storage
}

func NewHandlers(store storage.Storage) *Handlers {
	return &Handlers{store: store}
}

func getId(c *fiber.Ctx) (int, error) {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
