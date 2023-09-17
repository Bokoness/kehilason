package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	CommunitiesRoutes(app.Group("/communities"))
	AuthRoutes(app.Group("/auth"))
}
