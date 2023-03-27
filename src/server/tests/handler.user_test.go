package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/app/model"
	"github.com/HRemonen/kanban-board/setup"

	"github.com/stretchr/testify/assert"
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

	tests := []struct {
		description string

		route string

		expectedError  bool
		expectedStatus string
		expectedCode   int
		expectedBody   string
	}{
		{
			description:    "user index route",
			route:          "/api/v1/user",
			expectedError:  false,
			expectedStatus: "success",
			expectedCode:   200,
			expectedBody:   "",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		var body map[string]interface{}

		err = json.NewDecoder(res.Body).Decode(&body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		if err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}

		if body["status"] != test.expectedStatus {
			t.Errorf("Expected status '%s' but got '%s'", test.expectedStatus, body["status"])
		}

		data := body["data"].([]interface{})
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
}
