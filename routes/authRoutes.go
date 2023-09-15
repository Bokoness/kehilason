package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	route.Post("/register", handlers.RegisterUser)
	route.Post("/login", handlers.LoginUser)
}
