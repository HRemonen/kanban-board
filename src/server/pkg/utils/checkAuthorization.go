package utils

import (
	"errors"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// CheckAuthorization ... Check authorization of the user
// @Summary Check authorization of the user
// @Description Checks if the user sending request is actually the user
// @Description which holds the wanted resource.
// @Tags Utils
func IsAuthorized(c *fiber.Ctx, user model.User) (bool, error) {
	if user.ID == uuid.Nil {
		return false, errors.New("User not found")
	}

	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	authUserId := claims["sub"].(string)

	if user.ID.String() != authUserId {
		return false, errors.New("Unauthorized action")
	}

	return true, nil
}

// CheckAuthorization ... Check authorization of the user
// @Summary Check authorization of the user
// @Description Checks if the user sending request is actually the user
// @Description which holds the wanted resource.
// @Tags Utils
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
