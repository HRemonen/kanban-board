package router

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

func SwaggerRoute(app *fiber.App) {

	swag := app.Group("/swagger")

	swag.Get("*", swagger.HandlerDefault)
}
