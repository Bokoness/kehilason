package routes

import (
	"github.com/bokoness/kehilashon/api/handlers"
	"github.com/bokoness/kehilashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	route := app.Group("/auth")

	route.Post("/register", handlers.RegisterUser)
	route.Post("/login", handlers.LoginUser)
	route.Get("/check-auth", middleware.AuthMiddleware, handlers.CheckAuth)
}
