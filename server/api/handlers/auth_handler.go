package handlers

import (
	"errors"

	"github.com/bokoness/kehilashon/dto"
	"github.com/bokoness/kehilashon/models"
	"github.com/bokoness/kehilashon/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	var body models.User

	if err := services.ValidateRequestBody(c, new(dto.CreateUser), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := services.GetUserByEmail(body.Email)

	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "email already exists")
	}

	user, err := services.CreateUser(body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	return c.JSON(user.Clean())
}

func LoginUser(c *fiber.Ctx) error {
	var body models.User

	if err := services.ValidateRequestBody(c, new(dto.LoginUser), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	found, err := services.GetUserByEmail(body.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "פרטים לא נכונים")
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(body.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	jwt, err := services.GenerateAuthJWT(found.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = jwt

	c.Cookie(cookie)

	return c.JSON(found.Clean())
}

func CheckAuth(c *fiber.Ctx) error {
	if user, err := services.GetUserFromSession(c); err != nil {
		log.Error(err)

		return fiber.NewError(fiber.StatusUnauthorized)
	} else {
		return c.JSON(user.Clean())
	}
}
