package middleware

import (
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func CommunityMiddleware(c *fiber.Ctx) error {
	user, err := services.GetUserFromSession(c)

	if err != nil || user == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	isInCommunioty := services.UserIsInCommuinty(user, c.Params("id"))

	if !isInCommunioty {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
