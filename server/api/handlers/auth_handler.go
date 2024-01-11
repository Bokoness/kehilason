package handlers

import (
	"errors"

	"github.com/bokoness/lashon/dto"
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	userModel := new(models.User)

	if err := c.BodyParser(userModel); err != nil {
		return err
	}
	if err := services.ValidateRequestBody(c, new(dto.CreateUser), &userModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err := services.GetUserByEmail(userModel.Email)

	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "email already exists")
	}

	user, err := services.CreateUser(*userModel)

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
		return fiber.NewError(fiber.StatusUnauthorized)
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

func CheckLogin(c *fiber.Ctx) error {
	if user, err := services.GetUserFromSession(c); err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusUnauthorized)
	} else {
		return c.JSON(user.Clean())
	}
}
