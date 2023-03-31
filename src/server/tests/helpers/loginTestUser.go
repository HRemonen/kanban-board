package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ResponseBody struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    model.LoginData `json:"data"`
}

func LoginTestUser(app *fiber.App, db *gorm.DB) (*ResponseBody, error) {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/auth/login",
		strings.NewReader(`{"email": "alice@example.com", "password": "salainensalasana"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req, -1)

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)

	var body ResponseBody
	if err := json.Unmarshal(bytes, &body); err != nil {
		return &body, err
	}

	return &body, nil
}
