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

func TestGetAllBoards(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	db := database.DB.Db

	err := helpers.ClearTestBoards(db)
	if err != nil {
		t.Fatalf("Failed to clear test board entries: %v", err)
	}

	err = helpers.SeedTestBoards(db)
	if err != nil {
		t.Fatal("Failed to seed the test board entries", err)
	}

	tests := []struct {
		description string

		route string

		expectedStatus     string
		expectedCode       int
		expectedBody       string
		expectedBodyLength int
	}{
		{
			description:        "board index route",
			route:              "/api/v1/board",
			expectedStatus:     "success",
			expectedCode:       200,
			expectedBody:       "",
			expectedBodyLength: 5,
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

		assert.Equalf(t, test.expectedStatus, body["status"], "Test: '%s'. Expected status '%s' but got '%s'", test.description, test.expectedStatus, body["status"])
		assert.Equalf(t, test.expectedCode, res.StatusCode, "Test: '%s'. Expected HTTP statuscode '%s' but got '%s'", test.description, test.expectedCode, res.StatusCode)

		data := body["data"].([]interface{})
		assert.Equalf(t, test.expectedBodyLength, len(data), "Test: '%s'. Expected boards count '%s' but got '%s'", test.description, test.expectedBodyLength, len(data))
	}
}
