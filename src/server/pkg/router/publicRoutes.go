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
	board.Get("/:id", handlers.GetSingleBoard)

	// Board related list endpoints
	board.Post("/:id/list", handlers.CreateBoardList)
	board.Put("/:id/list/:list", handlers.UpdateBoardListPosition)
	board.Delete("/:id/list/:list", handlers.DeleteBoardList)

	list := v1.Group("/list")
	/* list.Get("/:id")
	list.Put("/:id") */

	// List related card enpoints
	list.Post("/:id/card", handlers.CreateListCard)
	list.Put("/:id/card/:card", handlers.UpdateListCardPosition)
	// list.Delete("/:id/card/:card")

	card := v1.Group("/card")

	card.Get("/:id", handlers.GetSingleCard)
	// card.Put("/:id")

}
