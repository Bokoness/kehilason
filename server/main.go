package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/routes"
	"github.com/bokoness/lashon/services"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDb()

	app := fiber.New()

	// Initialize default config
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))
	routes.Setup(app)

	gob.Register(&models.User{})

	services.SetupSessionStore()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
