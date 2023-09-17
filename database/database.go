package database

import (
	"fmt"
	"github.com/bokoness/lashon/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

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

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	log.Info("Connected to database")

	// auto migrations
	if err := db.AutoMigrate(&models.Community{}, &models.User{}, &models.CommunitiesUsers{}); err != nil {
		log.Error("Failed to migrate")
	}

	DB = db
}
