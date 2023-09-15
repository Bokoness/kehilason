package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	CampaignRoutes(app.Group("/campaigns"))
	AuthRoutes(app.Group("/auth"))
}
