package handlers

import (
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func GetCommunities(c *fiber.Ctx) error {
	return c.JSON(services.GetCommunities())
}

func GetCommunity(c *fiber.Ctx) error {
	return c.JSON(services.GetCommunity(c.Params("id")))
}

func CreateCommunity(c *fiber.Ctx) error {
	var body models.Community

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if found := services.GetCommunity(body.ID); found.ID != "" {
		return fiber.NewError(fiber.StatusNotAcceptable, "Community ID already exists")
	}

	return c.JSON(services.CreateCommunity(body))
}

func UpdateCommunity(c *fiber.Ctx) error {
	var body models.Community

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(services.UpdateCommunity(c.Params("id"), body))
}

func DeleteCommunity(c *fiber.Ctx) error {
	services.DeleteCommunity(c.Params("id"))

	return c.SendStatus(fiber.StatusOK)
}
