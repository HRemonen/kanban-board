package utils

import (
	"errors"

	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// IsAuthorized ... Check authorization of the user
// @Summary Check authorization of the user and return boolean and error if occured
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
