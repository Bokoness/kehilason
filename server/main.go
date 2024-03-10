package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/bokoness/kehilashon/database"
	"github.com/bokoness/kehilashon/models"
	"github.com/bokoness/kehilashon/routes"
	"github.com/bokoness/kehilashon/services"
	"github.com/joho/godotenv"

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
