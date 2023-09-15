package main

import (
	"encoding/gob"
	"fmt"
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/routes"
	"github.com/bokoness/lashon/services"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDb()

	app := fiber.New()

	routes.Setup(app)

	gob.Register(&models.User{})

	services.SetupSessionStore()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
