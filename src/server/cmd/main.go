package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

	"github.com/HRemonen/kanban-board/config"
	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/handlers"
	"github.com/HRemonen/kanban-board/router"
)

func main() {
	database.Connect()

	app := fiber.New()

	config.GoogleConfig()

	app.Use(logger.New())
	app.Use(cors.New())

	router.PublicRoutes(app)
	router.PrivateRoutes(app)

	app.Get("/login", handlers.GoogleLogin)
	app.Get("/login-callback", handlers.GoogleCallback)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	app.Listen(":8080")
}
