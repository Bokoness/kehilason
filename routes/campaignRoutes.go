package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/bokoness/lashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func CampaignRoutes(route fiber.Router) {
	route.Get("/:id", handlers.GetCampaign)
	route.Get("/", handlers.GetCampaigns)
	route.Post("/", middleware.AuthMiddleware, handlers.CreateCampaign)
	route.Put("/:id", middleware.AuthMiddleware, handlers.UpdateCampaign)
	route.Delete("/:id", middleware.AuthMiddleware, handlers.DeleteCampaign)
}
