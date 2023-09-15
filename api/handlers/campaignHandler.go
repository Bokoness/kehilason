package handlers

import (
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
)

func GetCampaigns(c *fiber.Ctx) error {
	return c.JSON(services.GetCampaigns())
}

func GetCampaign(c *fiber.Ctx) error {
	return c.JSON(services.GetCampaign(c.Params("id")))
}

func CreateCampaign(c *fiber.Ctx) error {
	var body models.Campaign

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if found := services.GetCampaign(body.ID); found.ID != "" {
		return fiber.NewError(fiber.StatusNotAcceptable, "Campaign ID already exists")
	}

	return c.JSON(services.CreateCampaign(body))
}

func UpdateCampaign(c *fiber.Ctx) error {
	var body models.Campaign

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(services.UpdateCampaign(c.Params("id"), body))
}

func DeleteCampaign(c *fiber.Ctx) error {
	services.DeleteCampaign(c.Params("id"))

	return c.SendStatus(fiber.StatusOK)
}
