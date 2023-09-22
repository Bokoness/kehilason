package handlers

import (
	"github.com/bokoness/lashon/dto"
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	var body dto.CreateUser

	if err := services.ValidateRequestBody(c, new(dto.CreateUser), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if found := services.GetUser(body.Email); found.Email != "" {
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

	found := services.GetUser(body.Email)

	if found.Email == "" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(body.Password))

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	jwt, err := services.GenerateAuthJWT(found.Email)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = jwt

	c.Cookie(cookie)

	return c.JSON(found.Clean())
}
