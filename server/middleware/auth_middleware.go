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

	// insert user to session
	statusCodeError := services.InsertUserToSession(c, user)
	if statusCodeError != 0 {
		return fiber.NewError(statusCodeError)
	}

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
