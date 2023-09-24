package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func SetupSessionStore() {
	Store = session.New()
}

func GetStore(c *fiber.Ctx) (*session.Session, error) {
	if sess, err := Store.Get(c); err != nil {
		return nil, err
	} else {
		return sess, nil
	}
}
