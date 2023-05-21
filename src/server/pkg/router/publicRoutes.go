package router

import (
	"github.com/HRemonen/kanban-board/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {

	// Google specific routes
	google := app.Group("/google")

	google.Get("/login", handlers.GoogleLogin)
	google.Get("/login-callback", handlers.GoogleCallback)

	// REST
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")

	auth.Post("/login", handlers.Login)

	user := v1.Group("/user")

	user.Get("/", handlers.GetAllUsers)
	user.Post("/", handlers.CreateUser)

	board := v1.Group("/board")

	board.Get("/", handlers.GetAllBoards)

	list := v1.Group("/list")

	list.Get(":id", handlers.GetSingleList)

	// List related card endpoints
	list.Post("/:id/card", handlers.CreateListCard)
	list.Put("/:id/card/:card", handlers.UpdateListCardPosition)
	list.Delete("/:id/card/:card", handlers.DeleteListCard)

	card := v1.Group("/card")

	card.Get("/:id", handlers.GetSingleCard)
}
