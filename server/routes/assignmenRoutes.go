package routes

import (
	"github.com/bokoness/lashon/api/handlers"
	"github.com/bokoness/lashon/middleware"
	"github.com/gofiber/fiber/v2"
)

func AssignmentRoutes(app fiber.Router) {
	route := app.Group("/communities/:communityId/assignments", middleware.AuthMiddleware)

	route.Get("/:id", handlers.GetAssignment)
	route.Get("/", handlers.GetAssignments)
	route.Post("/", handlers.CreateAssignment)
	route.Put("/:id", handlers.UpdateAssignment)
	route.Delete("/:id", handlers.DeleteAssignment)
}
