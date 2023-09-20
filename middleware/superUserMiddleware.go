package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SuperUserMiddleware(c *fiber.Ctx) error {
	user, err := GetUserFromCookiesByJWT(c)
	if err != nil || user == nil {
		return err
	}

	if !user.IsSuperuser {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// insert user to session
	statusCodeError := InsertUserToSession(c, user)
	if statusCodeError != 0 {
		return fiber.NewError(statusCodeError)
	}

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
