package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"fullName" gorm:"column:full_name"`
	IsSuperuser bool   `json:"isSuperuser" gorm:"column:is_superuser;default:false"`

	Communities []*Community `gorm:"many2many:communities_users"`
}

type CleanUser struct {
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

// BeforeCreate gorm hook that will hash the user's password before creation
func (u *User) BeforeCreate() (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("password hashing failed")
	}

	u.Password = string(hash)

	return
}

func (u User) Clean() CleanUser {
	return CleanUser{
		Email:    u.Email,
		FullName: u.FullName,
	}
}
