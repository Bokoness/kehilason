package services

import (
	"github.com/bokoness/kehilashon/database"
	"github.com/bokoness/kehilashon/models"
	"github.com/gofiber/fiber/v2/log"
)

func GetUser(uid uint) (*models.User, error) {
	var user models.User

	if err := database.DB.First(&user, uid).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &user, nil
}

func CreateUser(data models.User) (*models.User, error) {
	if err := database.DB.Create(&data).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &data, nil
}
