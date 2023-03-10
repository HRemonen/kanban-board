package utils

import (
	"errors"

	"github.com/HRemonen/kanban-board/database"
	"github.com/HRemonen/kanban-board/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CheckAuthorization(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db

	id := c.Params("id")

	var user model.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return user, errors.New("User not found")
	}

	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	authUserId := claims["sub"].(string)

	if user.ID.String() != authUserId {
		return user, errors.New("Unauthorized action")
	}

	return user, nil
}
