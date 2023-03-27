package main

import (
	_ "github.com/lib/pq"

	"github.com/HRemonen/kanban-board/app/config"
	"github.com/HRemonen/kanban-board/app/database"
	_ "github.com/HRemonen/kanban-board/docs"
	"github.com/HRemonen/kanban-board/setup"
)

// @title           Kanri API
// @description     Kanri is a Kanban board application
// @termsOfService  http://github.com/HRemonen/kanban-board/

// @contact.name   API Support
// @contact.url    http://github.com/HRemonen/kanban-board/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	database.Connect()

	app := setup.Setup()

	config.GoogleConfig()

	app.Listen(":8080")
}
