package router

import (
	"fmt"
	"github.com/godofprodev/tally/api/config"
	"github.com/godofprodev/tally/api/handlers"
	"github.com/godofprodev/tally/api/storage"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app   *fiber.App
	store *storage.MongoStore
}

func NewRouter(store *storage.MongoStore) *Router {
	app := fiber.New()

	return &Router{
		app:   app,
		store: store,
	}
}
func (r Router) RegisterHandlers() {
	h := handlers.NewHandlers(r.store)

	r.app.Get("/ping", h.HandlePing)
}

func (r Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
}
