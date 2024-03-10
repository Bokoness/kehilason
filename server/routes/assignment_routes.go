package routes

import (
	"github.com/bokoness/kehilashon/api/handlers"
	"github.com/bokoness/kehilashon/middleware"
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
