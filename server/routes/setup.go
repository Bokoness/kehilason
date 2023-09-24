package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	AuthRoutes(app)
	CommunitiesRoutes(app)
	AssignmentRoutes(app)
}
