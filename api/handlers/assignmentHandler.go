package handlers

import (
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAssignments(c *fiber.Ctx) error {
	return c.JSON(services.GetAssignments(c.Params("communityId")))
}

func GetAssignment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "assignmentId is not valid")
	}

	return c.JSON(services.GetAssignment(id))
}

func CreateAssignment(c *fiber.Ctx) error {
	var assignment models.Assignment

	if err := c.BodyParser(&assignment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	assignment.CommunityID = c.Params("communityId")

	return c.JSON(services.CreateAssignment(assignment))
}

func UpdateAssignment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "assignmentId is not valid")
	}

	var body models.Assignment

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(services.UpdateAssignment(id, body))
}

func DeleteAssignment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "assignmentId is not valid")
	}

	services.DeleteAssignment(id)

	return c.SendStatus(fiber.StatusOK)
}
