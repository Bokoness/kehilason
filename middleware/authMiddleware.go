package middleware

import (
	"fmt"
	"log"

	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func InsertUserToSession(c *fiber.Ctx, user *models.User) int {
	sess, err := services.GetStore(c)
	statusCodeError := fiber.StatusInternalServerError
	if err != nil {
		return statusCodeError
	}

	sess.Set("user", user)

	if err = sess.Save(); err != nil {
		log.Fatal(fmt.Printf("this is the error %s", err))
		return statusCodeError
	}

	return 0
}
func GetUserFromCookiesByJWT(c *fiber.Ctx) (*models.User, error) {
	cookie := c.Cookies("user")

	if cookie == "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized)
	}

	user, err := services.GetUserByJWT(cookie)
	return user, err
}
func AuthMiddleware(c *fiber.Ctx) error {
	user, err := GetUserFromCookiesByJWT(c)
	if err != nil || user == nil {
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
