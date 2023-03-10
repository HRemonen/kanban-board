package router

import (
	"github.com/HRemonen/kanban-board/handlers"
	"github.com/HRemonen/kanban-board/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")

	user.Get("/:id", middleware.JWTProtected(), handlers.GetSingleUser)
	user.Put("/:id", middleware.JWTProtected(), handlers.UpdateUser)
	user.Delete("/:id", middleware.JWTProtected(), handlers.DeleteUserByID)

	board := v1.Group("/board")

	board.Post("/", middleware.JWTProtected(), handlers.CreateBoard)
}
