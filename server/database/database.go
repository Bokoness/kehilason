package database

import (
	"fmt"
	"os"

	"github.com/bokoness/lashon/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// auto migrations
func mogation(db *gorm.DB) {
	dbModels := map[string]interface{}{
		"Community":        &models.Community{},
		"User":             &models.User{},
		"CommunitiesUsers": &models.CommunitiesUsers{},
		"Assignment":       &models.Assignment{},
		"Sessions":         &models.Sessions{},
	}
	for key, value := range dbModels {
		if err := db.AutoMigrate(value); err != nil {
			log.Error("Failed to migrate " + key)
		}
	}
}

func InitDb() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local", user, password, host, port)

	db, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Create the database if it doesn't exist
	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	db.Exec("USE " + dbname)

	log.Info("Connected to database " + dbname)

	mogation(db)

	DB = db
}
