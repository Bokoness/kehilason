package services

import (
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
)

func GetAssignments(assignmentId string) []models.Assignment {
	var assignments []models.Assignment

	database.DB.Find(&assignments, "communityId = ?", assignmentId)

	return assignments
}

func GetAssignment(id int) models.Assignment {
	var assignment models.Assignment

	database.DB.First(&assignment, "id = ?", id)

	return assignment
}

func CreateAssignment(data models.Assignment) models.Assignment {
	database.DB.Create(&data)

	return data
}

func UpdateAssignment(id int, data models.Assignment) models.Assignment {
	var result models.Assignment

	database.DB.Where("id = ?", id).First(&result)

	result.Start = data.Start
	result.End = data.End
	result.StartHour = data.StartHour
	result.EndHour = data.EndHour

	database.DB.Save(&result)

	return result
}

func DeleteAssignment(id int) {
	database.DB.Unscoped().Delete(&models.Assignment{}, id)
}
