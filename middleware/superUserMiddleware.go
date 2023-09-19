package middleware

import (
	"fmt"

	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func SuperUserMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("user")

	if cookie == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := services.GetUserByJWT(cookie)

	if !user.IsSuperuser {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// insert user to session
	sess, err := services.GetStore(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	sess.Set("user", user)

	if err = sess.Save(); err != nil {
		fmt.Printf("this is the error %s", err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
