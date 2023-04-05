package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/HRemonen/kanban-board/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	p := os.Getenv("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Helsinki", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{}, &model.Board{}, &model.List{}, &model.Card{})
	DB = Dbinstance{
		Db: db,
	}
}

func SetupTestDB() {
	host := "postgres"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the test database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		os.Exit(1)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running test database migrations")

	db.AutoMigrate(&model.User{}, &model.Board{}, &model.List{}, &model.Card{})
	DB = Dbinstance{
		Db: db,
	}
}
