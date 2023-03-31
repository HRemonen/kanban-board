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

	err := helpers.SeedTestUsers(db)
	if err != nil {
		t.Fatal("Failed to seed the test database", err)
	}

	defer helpers.ClearTestUsers(db)

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
		if len(data) != 3 {
			t.Errorf("Expected 3 users but got %d", len(data))
		}
	}
}

/* func TestGetSingleUser(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	tests := []struct {
		description string

		route string

		expectedError   bool
		expectedStatus  string
		expectedCode    int
		expectedMessage string
	}{
		{
			description:     "get user by ID when authenticated succeeds",
			route:           "/api/v1/user",
			expectedError:   false,
			expectedStatus:  "success",
			expectedCode:    200,
			expectedMessage: "User Found",
		},
		{
			description:     "get user by ID when not authenticated fails",
			route:           "/api/v1/user",
			expectedError:   true,
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

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	}
} */
