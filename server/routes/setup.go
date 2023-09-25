package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	route := app.Group("/api")
	AuthRoutes(route)
	CommunitiesRoutes(route)
	AssignmentRoutes(route)
}
