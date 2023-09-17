package services

import (
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
)

func GetCommunities() []models.Community {
	var communities []models.Community

	database.DB.Find(&communities)

	return communities
}

func GetCommunity(id string) models.Community {
	var community models.Community

	database.DB.First(&community, "id = ?", id)

	return community
}

func CreateCommunity(data models.Community) models.Community {
	database.DB.Create(&data)

	return data
}

func UpdateCommunity(id string, data models.Community) models.Community {
	var result models.Community

	database.DB.Where("id = ?", id).First(&result)

	result.Name = data.Name

	database.DB.Save(&result)

	return result
}

func DeleteCommunity(id string) {
	record := models.Community{ID: id}
	database.DB.Unscoped().Delete(&record)
}
