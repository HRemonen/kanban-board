package handlers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/HRemonen/kanban-board/app/config"
	"github.com/gofiber/fiber/v2"
)

// GoogleLogin ... Google OAuth login
// @Summary Google OAuth login
// @Description Google OAuth login
// @Tags Login
// @Router /google/login [post]
func GoogleLogin(c *fiber.Ctx) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL(config.Config("RANDOM_STATE"))

	c.Status(303)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != config.Config("RANDOM_STATE") {
		return c.SendString("States does not match")
	}

	code := c.Query("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(userData))

}
