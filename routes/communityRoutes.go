package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/bokoness/lashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommunitiesRoutes(route fiber.Router) {
	route.Get("/:id", handlers.GetCommunity)
	route.Get("/", handlers.GetCommunities)
	route.Post("/", middleware.SuperUserMiddleware, handlers.CreateCommunity)
	route.Put("/:id", middleware.AuthMiddleware, handlers.UpdateCommunity)
	route.Delete("/:id", middleware.SuperUserMiddleware, handlers.DeleteCommunity)
}
