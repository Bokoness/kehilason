package routes

import (
	"github.com/bokoness/kehilashon/api/handlers"
	"github.com/bokoness/kehilashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommunitiesRoutes(app fiber.Router) {
	route := app.Group("/communities")

	route.Get("/:id", handlers.GetCommunity)
	route.Get("/", handlers.GetCommunities)
	route.Post("/", middleware.SuperUserMiddleware, handlers.CreateCommunity)
	route.Put("/:id", middleware.AuthMiddleware, handlers.UpdateCommunity)
	route.Delete("/:id", middleware.SuperUserMiddleware, handlers.DeleteCommunity)
	route.Post("/:id/register-user", middleware.AuthMiddleware, handlers.RegisterUserToCommunity)
}
