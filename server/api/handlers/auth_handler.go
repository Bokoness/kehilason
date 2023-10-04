package handlers

import (
	"github.com/bokoness/lashon/dto"
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	var body models.User

	if err := services.ValidateRequestBody(c, new(dto.CreateUser), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if found := services.GetUserByEmail(body.Email); found.Email != "" {
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

	found := services.GetUserByEmail(body.Email)

	if found.Email == "" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(body.Password))

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
func RefreshToken(c *fiber.Ctx) error {
	user, err := services.GetUserFromSession(c)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	userID := user.ID // Replace with your user ID logic
	token, err := services.GenerateRefreshToken(userID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
