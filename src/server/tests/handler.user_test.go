package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/HRemonen/kanban-board/app/database"
	"github.com/HRemonen/kanban-board/setup"
	"github.com/HRemonen/kanban-board/tests/helpers"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	db := database.DB.Db

	err := helpers.ClearTestUsers(db)
	if err != nil {
		t.Fatalf("Failed to clear test user entries: %v", err)
	}

	err = helpers.SeedTestUsers(db)
	if err != nil {
		t.Fatal("Failed to seed the test user entries", err)
	}

	tests := []struct {
		description string

		route string

		expectedStatus string
		expectedCode   int
		expectedBody   string
	}{
		{
			description:    "returns all users",
			route:          "/api/v1/user",
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
		if len(data) != 3 {
			t.Errorf("Expected 3 users but got %d", len(data))
		}
	}
}

func TestGetSingleUser(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	db := database.DB.Db

	err := helpers.ClearTestUsers(db)
	if err != nil {
		t.Fatalf("Failed to clear test data: %v", err)
	}

	err = helpers.SeedTestUsers(db)
	if err != nil {
		t.Fatal("Failed to seed the test database", err)
	}

	user, err := helpers.LoginTestUser(app, db)

	tests := []struct {
		description string

		route string

		token string

		expectedStatus  string
		expectedCode    int
		expectedMessage string
	}{
		{
			description:     "returns user by ID when authenticated",
			route:           "/api/v1/user/" + user.Data.User.ID.String(),
			token:           user.Data.Token,
			expectedStatus:  "success",
			expectedCode:    200,
			expectedMessage: "User Found",
		},
		{
			description:     "does not return user by ID when not authenticated",
			route:           "/api/v1/user/" + user.Data.User.ID.String(),
			token:           "",
			expectedStatus:  "error",
			expectedCode:    400,
			expectedMessage: "Missing or malformed JWT",
		},
		{
			description:     "does not return user by ID when unauthorized",
			route:           "/api/v1/user/" + user.Data.User.ID.String(),
			token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA0NjE4MjYsImlhdCI6MTY4MDQyNTgyNiwibmJmIjoxNjgwNDI1ODI2LCJzdWIiOiIyN2MyM2ViNi04MThiLTRlYTMtOWU1MC04MjAwMDFkYTY0NWUifQ.k1irIqJ93ACScqVcBkXPHpS8dZTpCc2V7LFZPb-KBKw",
			expectedStatus:  "error",
			expectedCode:    401,
			expectedMessage: "Unauthorized action",
		},
		{
			description:     "does not return non existing user by ID",
			route:           "/api/v1/user/c33234a3-8f08-489e-91a9-e03c9c167a64", // This User ID does not exist
			token:           user.Data.Token,
			expectedStatus:  "error",
			expectedCode:    404,
			expectedMessage: "Could not fetch user",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(
			"GET",
			test.route,
			nil,
		)
		bearer := "Bearer " + test.token
		req.Header.Set("Authorization", bearer)

		res, _ := app.Test(req, -1)

		var body map[string]interface{}

		err = json.NewDecoder(res.Body).Decode(&body)

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		assert.Equalf(t, test.expectedStatus, body["status"], test.description)
		assert.Equalf(t, test.expectedMessage, body["message"], test.description)
	}
}
