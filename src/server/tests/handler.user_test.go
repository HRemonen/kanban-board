package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetAllUsers(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	db := database.DB.Db

	// Insert some test users into the database
	user1 := model.User{Name: "Alice", Email: "alice@example.com", Password: "salainensalasana"}
	user2 := model.User{Name: "Bob", Email: "bob@example.com", Password: "salainensalasana"}

	db.Create(&user1)
	db.Create(&user2)

	req := httptest.NewRequest(fiber.MethodGet, "/api/v1/user", nil)

	resp, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
