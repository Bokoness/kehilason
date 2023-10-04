package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"fullName" gorm:"column:full_name"`
	IsSuperuser bool   `json:"isSuperuser" gorm:""column:is_superuser;default:false"`

	Communities []*Community `gorm:"many2many:communities_users"`
}

type CleanUser struct {
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

func (u User) Clean() CleanUser {
	return CleanUser{
		Email:    u.Email,
		FullName: u.FullName,
	}
}
