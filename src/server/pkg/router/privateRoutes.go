package router

import (
	"github.com/HRemonen/kanban-board/app/handlers"
	"github.com/HRemonen/kanban-board/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")

	user.Get("/:id", middleware.JWTProtected(), handlers.GetSingleUser)
	user.Put("/:id", middleware.JWTProtected(), handlers.UpdateUser)
	user.Delete("/:id", middleware.JWTProtected(), handlers.DeleteUser)

	board := v1.Group("/board")

	board.Get("/:id", middleware.JWTProtected(), handlers.GetSingleBoard)
	board.Get("/:userID/boards", middleware.JWTProtected(), handlers.GetUserBoards)
	board.Post("/", middleware.JWTProtected(), handlers.CreateBoard)
	board.Delete("/:id", middleware.JWTProtected(), handlers.DeleteBoard)

	// Board related list endpoints
	board.Post("/:id/list", middleware.JWTProtected(), handlers.CreateBoardList)
	board.Put("/:id/list/:list", middleware.JWTProtected(), handlers.UpdateBoardListPosition)
	board.Delete("/:id/list/:list", middleware.JWTProtected(), handlers.DeleteBoardList)
}
