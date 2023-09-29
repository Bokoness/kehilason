package services

import (
	"errors"
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
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

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	data.Password = string(hash)

	if err := database.DB.Create(&data).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	return &data, nil
}
