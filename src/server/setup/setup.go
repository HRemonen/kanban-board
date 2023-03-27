package setup

import (
	_ "github.com/HRemonen/kanban-board/docs"
	"github.com/HRemonen/kanban-board/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func Setup() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	router.PublicRoutes(app)
	router.PrivateRoutes(app)
	router.SwaggerRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app
}
