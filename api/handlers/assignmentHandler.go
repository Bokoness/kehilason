package handlers

import (
	"github.com/bokoness/lashon/dto"
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
	var body models.Assignment

	if err := services.ValidateRequestBody(c, new(dto.CreateUpdateAssignment), &body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body.CommunityID = c.Params("communityId")

	user, err := services.GetUserFromSession(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	body.UserID = user.Email

	return c.JSON(services.CreateAssignment(body))
}

func UpdateAssignment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "assignmentId is not valid")
	}

	var body models.Assignment
	if err := services.ValidateRequestBody(c, new(dto.CreateUpdateAssignment), &body); err != nil {
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
