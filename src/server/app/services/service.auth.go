package services

import (
	"errors"
	"strings"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) (model.LoginData, error) {
	db := database.DB.Db
	payload := new(model.LoginUserInput)

	c.BodyParser(payload)

	var user model.User
	var loginData model.LoginData
	var validate = utils.NewValidator()

	err := validate.Struct(payload)

	if err != nil {
		return loginData, err
	}

	err = db.Model(&model.User{}).Preload("Boards").First(&user, "email = ?", strings.ToLower(payload.Email)).Error

	if err != nil {
		return loginData, err
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return loginData, errors.New("Invalid password")
	}

	if user.Provider == "Google" {
		return loginData, errors.New("Use Oauth login for Google login")
	}

	token, err := utils.GenerateNewAccessToken(user.ID)

	loginData.Token = token
	loginData.User = model.FilteredResponse(&user)

	return loginData, err
}
