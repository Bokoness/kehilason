package middleware

import (
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("user")

	if cookie == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := services.GetUserByJWT(cookie)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	err = services.SaveUserInSession(c, *user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.Next()
}
