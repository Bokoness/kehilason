package handlers

import (
	"errors"
	"github.com/bokoness/lashon/dto"
	"github.com/bokoness/lashon/models"
	"github.com/bokoness/lashon/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"strconv"
)

func GetAssignments(c *fiber.Ctx) error {
	assignments, err := services.GetAssignments(c.Params("communityId"))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(assignments)
}

func GetAssignment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "assignmentId is not valid")
	}

	assignment, err := services.GetAssignment(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound)
		} else {
			log.Error(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}
	}

	return c.JSON(assignment)
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

	body.UserID = user.ID

	assignment, err := services.CreateAssignment(body)

	if err != nil {
		log.Error(err)

		if services.IsDuplicatedKeyError(err) {
			return fiber.NewError(fiber.StatusNotAcceptable, "Assignment already exists")
		}

		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(assignment)
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

	assignment, err := services.UpdateAssignment(id, body)

	if err != nil {
		log.Error(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound)
		}

		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(assignment)
}

func DeleteAssignment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "assignmentId is not valid")
	}

	services.DeleteAssignment(id)

	return c.SendStatus(fiber.StatusOK)
}
