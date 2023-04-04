package tests

import (
	"encoding/json"
	"fmt"
	"io"
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
		t.Fatalf("Failed to clear test data: %v", err)
	}

	err = helpers.SeedTestUsers(db)
	if err != nil {
		t.Fatal("Failed to seed the test database", err)
	}

	tests := []struct {
		description string

		route string

		expectedStatus string
		expectedCode   int
		expectedBody   string
	}{
		{
			description:    "user index route",
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
			description:     "get user by ID when authenticated succeeds",
			route:           "/api/v1/user/" + user.Data.User.ID.String(),
			token:           user.Data.Token,
			expectedStatus:  "success",
			expectedCode:    200,
			expectedMessage: "User Found",
		},
		{
			description:     "get user by ID when not authenticated fails",
			route:           "/api/v1/user/" + user.Data.User.ID.String(),
			token:           "",
			expectedStatus:  "error",
			expectedCode:    400,
			expectedMessage: "Missing or malformed JWT",
		},
		{
			description:     "get user by ID when unauthorized fails",
			route:           "/api/v1/user/" + user.Data.User.ID.String(),
			token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA0NjE4MjYsImlhdCI6MTY4MDQyNTgyNiwibmJmIjoxNjgwNDI1ODI2LCJzdWIiOiIyN2MyM2ViNi04MThiLTRlYTMtOWU1MC04MjAwMDFkYTY0NWUifQ.k1irIqJ93ACScqVcBkXPHpS8dZTpCc2V7LFZPb-KBKw",
			expectedStatus:  "error",
			expectedCode:    401,
			expectedMessage: "Unauthorized action",
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

		b, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(b))

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	}
}
