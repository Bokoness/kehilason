package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/bokoness/lashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	route := app.Group("/authentication")

	route.Post("/register", handlers.RegisterUser)
	route.Post("/login", handlers.LoginUser)
	route.Get("/check-login", middleware.AuthMiddleware, handlers.CheckLogin)
}
