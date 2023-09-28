package services

import (
	"errors"
	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(uid uint) models.User {
	var user models.User

	database.DB.First(&user, uid)

	return user
}

func GetUserByEmail(email string) models.User {
	var user models.User

	database.DB.First(&user, "email = ?", email)

	return user
}

func CreateUser(data models.User) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	data.Password = string(hash)

	database.DB.Create(&data)

	return &data, nil
}
