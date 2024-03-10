package handlers

import (
	"errors"

	"github.com/bokoness/kehilashon/dto"
	"github.com/bokoness/kehilashon/models"
	"github.com/bokoness/kehilashon/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func GetCommunities(c *fiber.Ctx) error {
	communities, err := services.GetCommunities()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(communities)
}

func GetCommunity(c *fiber.Ctx) error {
	community, err := services.GetCommunity(c.Params("id"))
	if err != nil {
		log.Error(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound)
		}

		return fiber.NewError(fiber.StatusInternalServerError)

	}

	return c.JSON(community)
}

func CreateCommunity(c *fiber.Ctx) error {
	var body models.Community

	if err := services.ValidateRequestBody(c, new(dto.CreateCommunity), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	community, err := services.CreateCommunity(body)
	if err != nil {
		log.Error(err)

		if services.IsDuplicatedKeyError(err) {
			return fiber.NewError(fiber.StatusNotAcceptable, "Community ID already exists")
		}

		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(community)
}

func UpdateCommunity(c *fiber.Ctx) error {
	var body models.Community

	if err := services.ValidateRequestBody(c, new(dto.UpdateCommunity), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	community, err := services.UpdateCommunity(c.Params("id"), body)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.JSON(community)
}

func DeleteCommunity(c *fiber.Ctx) error {
	services.DeleteCommunity(c.Params("id"))

	return c.SendStatus(fiber.StatusOK)
}

func RegisterUserToCommunity(c *fiber.Ctx) error {
	user, err := services.GetUserFromSession(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	if err = services.RegisterUserToCommunity(user.ID, c.Params("id")); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
