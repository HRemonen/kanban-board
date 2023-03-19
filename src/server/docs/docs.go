// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://github.com/HRemonen/kanban-board/",
        "contact": {
            "name": "API Support",
            "url": "http://github.com/HRemonen/kanban-board/"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "login user and generate JWT token",
                "tags": [
                    "Login"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login attributes",
                        "name": "login_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginUserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.LoginData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/board": {
            "get": {
                "description": "get all boards",
                "tags": [
                    "Boards"
                ],
                "summary": "Get all boards",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.APIBoard"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new board",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Create a new board",
                "parameters": [
                    {
                        "description": "Board attributes",
                        "name": "board_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BoardUserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.APIBoard"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/board/{id}": {
            "get": {
                "description": "get a single board",
                "tags": [
                    "Boards"
                ],
                "summary": "Get a single board",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIBoard"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a board by ID",
                "tags": [
                    "Boards"
                ],
                "summary": "Delete a board by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/board/{id}/list": {
            "post": {
                "description": "create a new list for a board",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Create a new list for a board",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "List attributes",
                        "name": "list_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListUserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.List"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/board/{id}/list/{list}": {
            "put": {
                "description": "update list position on the board",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Update list position on the board",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "List position",
                        "name": "position",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListPositionInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a list from board",
                "tags": [
                    "Boards"
                ],
                "summary": "Delete a list from board",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "list",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/card/{id}": {
            "get": {
                "description": "get a single card by ID",
                "tags": [
                    "Cards"
                ],
                "summary": "Get a single card by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Card"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/google/login": {
            "post": {
                "description": "Google OAuth login",
                "tags": [
                    "Login"
                ],
                "summary": "Google OAuth login",
                "responses": {}
            }
        },
        "/list/{id}/card": {
            "post": {
                "description": "create a new card for a list",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Create a new card for a list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Card attributes",
                        "name": "card_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CardUserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Card"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/list/{id}/card/{card}": {
            "put": {
                "description": "update card position on the list",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Update card position on the list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "card",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Card position",
                        "name": "position",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CardPositionInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/list/{id}/card/{list}": {
            "delete": {
                "description": "delete a card from list",
                "tags": [
                    "Cards"
                ],
                "summary": "Delete a card from list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "card ID",
                        "name": "card",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "get all users",
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create new user based on paramters",
                "parameters": [
                    {
                        "description": "User attributes",
                        "name": "user_attrs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterUserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "get a single user by ID",
                "tags": [
                    "Users"
                ],
                "summary": "Get a single user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "update user name by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user name by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user by ID",
                "tags": [
                    "Users"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.APIBoard": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.List"
                    }
                },
                "name": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "model.Board": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.List"
                    }
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "model.BoardUserInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 100
                },
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "model.Card": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "listID": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.CardPositionInput": {
            "type": "object",
            "properties": {
                "position": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "model.CardUserInput": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "model.List": {
            "type": "object",
            "properties": {
                "boardID": {
                    "type": "string"
                },
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Card"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.ListPositionInput": {
            "type": "object",
            "properties": {
                "position": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "model.ListUserInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 1
                }
            }
        },
        "model.LoginData": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.UserResponse"
                }
            }
        },
        "model.LoginUserInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.RegisterUserInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "passwordConfirm"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "passwordConfirm": {
                    "type": "string"
                }
            }
        },
        "model.UpdateUser": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "boards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Board"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "boards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Board"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Kanri API",
	Description:      "Kanri is a Kanban board application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
