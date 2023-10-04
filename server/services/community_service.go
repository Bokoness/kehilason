package services

import (
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"github.com/gofiber/fiber/v2/log"
)

func GetCommunities() (*[]models.Community, error) {
	var communities []models.Community

	if err := database.DB.Find(&communities).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &communities, nil
}

func GetCommunity(id string) (*models.Community, error) {
	var community models.Community

	if err := database.DB.First(&community, "id = ?", id).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &community, nil
}

func CreateCommunity(data models.Community) (*models.Community, error) {
	if err := database.DB.Create(&data).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &data, nil
}

func UpdateCommunity(id string, data models.Community) (*models.Community, error) {
	var result models.Community

	if err := database.DB.Where("id = ?", id).First(&result).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	result.Name = data.Name

	database.DB.Save(&result)

	return &result, nil
}

func DeleteCommunity(id string) {
	record := models.Community{ID: id}
	database.DB.Unscoped().Delete(&record)
}

func RegisterUserToCommunity(userId uint, communityId string) error {
	record := models.CommunitiesUsers{CommunityID: communityId, UserID: userId}
	return database.DB.Save(&record).Error
}
