package utils

import (
	"errors"

	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// ExtractUser ... Extract the user from the request
// @Summary Extract the user from the request
// @Description Checks if the request.user is in the database
// @Description and returns the user or error.
// @Tags Utils
func ExtractUser(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db

	var user model.User

	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	authUserId := claims["sub"].(string)

	db.Find(&user, "id = ?", authUserId)

	if user.ID == uuid.Nil {
		return user, errors.New("User not found")
	}

	return user, nil
}
