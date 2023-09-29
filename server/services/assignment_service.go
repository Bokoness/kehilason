package services

import (
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"github.com/gofiber/fiber/v2/log"
)

func GetAssignments(assignmentId string) ([]*models.Assignment, error) {
	var assignments []*models.Assignment

	if err := database.DB.Find(&assignments, "community_id = ?", assignmentId).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return assignments, nil
}

func GetAssignment(id int) (*models.Assignment, error) {
	var assignment models.Assignment

	if err := database.DB.First(&assignment, "id = ?", id).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &assignment, nil
}

func CreateAssignment(data models.Assignment) (*models.Assignment, error) {
	if err := database.DB.Create(&data).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &data, nil

}

func UpdateAssignment(id int, data models.Assignment) (*models.Assignment, error) {
	var result models.Assignment

	if err := database.DB.Where("id = ?", id).First(&result).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	result.Start = data.Start
	result.End = data.End
	result.StartHour = data.StartHour
	result.EndHour = data.EndHour

	database.DB.Save(&result)

	return &result, nil
}

func DeleteAssignment(id int) {
	database.DB.Unscoped().Delete(&models.Assignment{}, id)
}
