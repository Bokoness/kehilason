package middleware

import (
	"github.com/bokoness/kehilashon/services"
	"github.com/gofiber/fiber/v2"
)

func SuperUserMiddleware(c *fiber.Ctx) error {
	user, err := services.GetUserFromCookiesByJWT(c)
	if err != nil || user == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !user.IsSuperuser {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = services.InsertUserToSession(c, user)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Next()
}
