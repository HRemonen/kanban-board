package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
)

func SwaggerRoute(app *fiber.App) {

	swag := app.Group("/swagger")

	swag.Get("*", swagger.HandlerDefault)
}
