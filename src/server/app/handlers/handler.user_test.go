package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/gofiber/fiber/v2"
)

func TestGetAllUsers(t *testing.T) {
	app := fiber.New()
	database.SetupTestDB() // Set up a test database

	// Insert some test users into the database
	db := database.DB.Db
	user1 := model.User{Name: "Alice", Email: "alice@example.com"}
	user2 := model.User{Name: "Bob", Email: "bob@example.com"}
	db.Create(&user1)
	db.Create(&user2)

	req := httptest.NewRequest("GET", "/api/v1/user", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200 but got %d", resp.StatusCode)
	}

	var respBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if respBody["status"] != "success" {
		t.Errorf("Expected status 'success' but got '%s'", respBody["status"])
	}

	data := respBody["data"].([]interface{})
	if len(data) != 2 {
		t.Errorf("Expected 2 users but got %d", len(data))
	}

	user1Data := data[0].(map[string]interface{})
	if user1Data["name"] != user1.Name || user1Data["email"] != user1.Email {
		t.Errorf("Unexpected user data: %v", user1Data)
	}

	user2Data := data[1].(map[string]interface{})
	if user2Data["name"] != user2.Name || user2Data["email"] != user2.Email {
		t.Errorf("Unexpected user data: %v", user2Data)
	}
}
