package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
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

	err := helpers.ClearTestUsers(db)
	err = helpers.ClearTestBoards(db)
	if err != nil {
		t.Fatalf("Failed to clear test data: %v", err)
	}

	err = helpers.SeedTestUsers(db)
	err = helpers.SeedTestBoards(db)
	if err != nil {
		t.Fatal("Failed to seed the test database", err)
	}

	tests := []struct {
		description string

		route string

		expectedStatus  string
		expectedCode    int
		expectedBody    string
		expectedMessage string
	}{
		{
			description:     "board index route",
			route:           "/api/v1/board",
			expectedStatus:  "success",
			expectedCode:    200,
			expectedBody:    "",
			expectedMessage: "Boards found",
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

		if test.expectedMessage != "" {
			assert.Equalf(t, body["message"], test.expectedMessage, "Test: '%s'. Expected return message '%s' but got '%s'", test.description, test.expectedMessage, body["message"])
		}
	}
}

func TestGetSingleBoard(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	db := database.DB.Db

	err := helpers.ClearTestUsers(db)
	err = helpers.ClearTestBoards(db)
	if err != nil {
		t.Fatalf("Failed to clear test data: %v", err)
	}

	err = helpers.SeedTestUsers(db)
	err = helpers.SeedTestBoards(db)
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
		expectedBody    string
		expectedMessage string
	}{
		{
			description:     "returns user's board by ID when authenticated",
			route:           "/api/v1/board/7a02f5d7-75aa-46b7-a698-c073ce49b12f", // Alice's (our test user) board
			token:           user.Data.Token,
			expectedStatus:  "success",
			expectedCode:    200,
			expectedBody:    "7a02f5d7-75aa-46b7-a698-c073ce49b12f",
			expectedMessage: "Board found",
		},
		{
			description:     "does not return board by ID when not authenticated",
			route:           "/api/v1/board/7a02f5d7-75aa-46b7-a698-c073ce49b12f", // Alice's (our test user) board
			token:           "",
			expectedStatus:  "error",
			expectedCode:    400,
			expectedBody:    "",
			expectedMessage: "Missing or malformed JWT",
		},
		{
			description:     "does not return board by ID when unauthorized",
			route:           "/api/v1/board/7a02f5d7-75aa-46b7-a698-c073ce49b12f", // Alice's (our test user) board
			token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA0NjE4MjYsImlhdCI6MTY4MDQyNTgyNiwibmJmIjoxNjgwNDI1ODI2LCJzdWIiOiIyN2MyM2ViNi04MThiLTRlYTMtOWU1MC04MjAwMDFkYTY0NWUifQ.k1irIqJ93ACScqVcBkXPHpS8dZTpCc2V7LFZPb-KBKw",
			expectedStatus:  "error",
			expectedCode:    401,
			expectedBody:    "",
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

		if test.expectedMessage != "" {
			assert.Equalf(t, body["message"], test.expectedMessage, "Test: '%s'. Expected return message '%s' but got '%s'", test.description, test.expectedMessage, body["message"])
		}
	}
}

func TestCreateBoard(t *testing.T) {
	database.SetupTestDB()
	app := setup.Setup()

	db := database.DB.Db

	err := helpers.ClearTestUsers(db)
	err = helpers.ClearTestBoards(db)
	if err != nil {
		t.Fatalf("Failed to clear test data: %v", err)
	}

	err = helpers.SeedTestUsers(db)
	err = helpers.SeedTestBoards(db)
	if err != nil {
		t.Fatal("Failed to seed the test database", err)
	}

	user, err := helpers.LoginTestUser(app, db)

	tests := []struct {
		description string

		route string

		token string

		body string

		expectedStatus  string
		expectedCode    int
		expectedBody    string
		expectedMessage string
	}{
		{
			description:     "creates a new board when authenticated",
			route:           "/api/v1/board",
			token:           user.Data.Token,
			body:            `{"name": "Test board", "description": "Testing create board handler"}`,
			expectedStatus:  "success",
			expectedCode:    201,
			expectedBody:    "",
			expectedMessage: "Board has been created",
		},
		{
			description:     "does not create board when not authenticated",
			route:           "/api/v1/board",
			token:           "",
			body:            `{"name": "Test board", "description": "Testing create board handler"}`,
			expectedStatus:  "error",
			expectedCode:    400,
			expectedBody:    "",
			expectedMessage: "Missing or malformed JWT",
		},
		{
			description:     "does not create board when unauthorized",
			route:           "/api/v1/board",
			token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA0NjE4MjYsImlhdCI6MTY4MDQyNTgyNiwibmJmIjoxNjgwNDI1ODI2LCJzdWIiOiIyN2MyM2ViNi04MThiLTRlYTMtOWU1MC04MjAwMDFkYTY0NWUifQ.k1irIqJ93ACScqVcBkXPHpS8dZTpCc2V7LFZPb-KBKw",
			body:            `{"name": "Test board", "description": "Testing create board handler"}`,
			expectedStatus:  "error",
			expectedCode:    401,
			expectedBody:    "",
			expectedMessage: "Unauthorized action",
		},
		{
			description:     "does not create board when missing name in payload",
			route:           "/api/v1/board",
			token:           user.Data.Token,
			body:            `{"description": "Testing create board handler"}`,
			expectedStatus:  "error",
			expectedCode:    422,
			expectedBody:    "",
			expectedMessage: "Validation of the inputs failed",
		},
		{
			description:     "does not create board when too short name in payload",
			route:           "/api/v1/board",
			token:           user.Data.Token,
			body:            `{"name": "a"}`,
			expectedStatus:  "error",
			expectedCode:    422,
			expectedBody:    "",
			expectedMessage: "Validation of the inputs failed",
		},
		{
			description:     "does not create board when too long name in payload",
			route:           "/api/v1/board",
			token:           user.Data.Token,
			body:            `{"name": "FbM3q8JjxRa1d2y4IxHckWref3qrOf2TIWiM6cbCHG2E7bjOpQ9t150HlIS2dcIMLhebqOrfijAAIzIfTrTCfM1DmwHTKRlitGCq9UkBmvTRW5t7rZ0BUjN0XMs2IQ5GAkLDOxBEIB1ODtWpspCETVXfrjqUot4vb7CRJsYFjr4j5wIDhkyNf5lcQGrNxgG8JY8u8vq2g4oGfTQ9itKUh70GTPyO486L0H1XMWn1xmoyvudsPOf0MswYjjQJ16R5"}`,
			expectedStatus:  "error",
			expectedCode:    422,
			expectedBody:    "",
			expectedMessage: "Validation of the inputs failed",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(
			"POST",
			test.route,
			strings.NewReader(test.body),
		)
		req.Header.Set("Content-Type", "application/json")
		bearer := "Bearer " + test.token
		req.Header.Set("Authorization", bearer)

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

		if test.expectedMessage != "" {
			assert.Equalf(t, body["message"], test.expectedMessage, "Test: '%s'. Expected return message '%s' but got '%s'", test.description, test.expectedMessage, body["message"])
		}
	}
}
