package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	route := app.Group("/auth")

	route.Post("/register", handlers.RegisterUser)
	route.Post("/login", handlers.LoginUser)
}
