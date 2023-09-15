package services

import (
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
)

func GetCampaigns() []models.Campaign {
	var campaigns []models.Campaign

	database.DB.Find(&campaigns)

	return campaigns
}

func GetCampaign(id string) models.Campaign {
	var campaign models.Campaign

	database.DB.First(&campaign, "id = ?", id)

	return campaign
}

func CreateCampaign(data models.Campaign) models.Campaign {
	database.DB.Create(&data)

	return data
}

func UpdateCampaign(id string, data models.Campaign) models.Campaign {
	var result models.Campaign

	database.DB.Where("id = ?", id).First(&result)

	result.Name = data.Name

	database.DB.Save(&result)

	return result
}

func DeleteCampaign(id string) {
	record := models.Campaign{ID: id}
	database.DB.Unscoped().Delete(&record)
}
