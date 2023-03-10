package router

import (
	"github.com/HRemonen/kanban-board/handlers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")

	user.Get("/", handlers.GetAllUsers)
	user.Post("/", handlers.CreateUser)

	auth := v1.Group("/auth")

	auth.Post("/login", handlers.Login)
}
