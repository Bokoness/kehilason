package services

import (
	"errors"
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/dto"
	"github.com/bokoness/lashon/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(email string) models.User {
	var user models.User

	database.DB.First(&user, "email = ?", email)

	return user
}

func CreateUser(dto dto.CreateUser) (*models.User, error) {
	var result models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	result.Email = dto.Email
	result.Password = string(hash)
	result.FullName = dto.FullName

	database.DB.Create(&result)

	return &result, nil
}
