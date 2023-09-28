package router

import (
	"fmt"
	"github.com/godofprodev/tally/api/config"
	"github.com/godofprodev/tally/api/handlers"
	"github.com/godofprodev/tally/api/responses"
	"github.com/godofprodev/tally/api/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Router struct {
	app   *fiber.App
	store storage.Storage
}

func NewRouter(store storage.Storage) *Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	return &Router{
		app:   app,
		store: store,
	}
}

func (r Router) RegisterMiddlewares() {
	r.app.Use(logger.New())
	r.app.Use(cors.New())
}

func (r Router) RegisterHandlers() {
	h := handlers.NewHandlers(r.store)

	v1 := r.app.Group("/v1")
	v1.Get("/monitor", monitor.New())

	v1.Get("/ping", h.HandlePing)
	v1.Get("/guilds", h.HandleGetGuilds)
	v1.Get("/guilds/:id", h.HandleGetGuild)
	v1.Get("/users/", h.HandleGetUsers)

	v1.Post("/guilds", h.HandleCreateGuild)
}

func (r Router) Listen(s *config.ServerConfig) error {
	return r.app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case responses.APIError:
		return c.Status(e.Status).JSON(e)
	case responses.APISuccessData:
		return c.Status(e.Status).JSON(e.Data)
	case responses.APISuccessResponse:
		return c.Status(e.Status).JSON(e)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"message": "internal server error"})
	}
}
