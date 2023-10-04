package middleware

import (
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	user, err := services.GetUserFromCookiesByJWT(c)
	if err != nil || user == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = services.InsertUserToSession(c, user)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
