package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

	"github.com/HRemonen/kanban-board/app/config"
	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/pkg/router"
)

// @title           Kanri API
// @description     Kanri is a Kanban board application
// @termsOfService  http://github.com/HRemonen/kanban-board/

// @contact.name   API Support
// @contact.url    http://github.com/HRemonen/kanban-board/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	database.Connect()

	app := fiber.New()

	config.GoogleConfig()

	app.Use(logger.New())
	app.Use(cors.New())

	router.PublicRoutes(app)
	router.PrivateRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(":8080")
}
