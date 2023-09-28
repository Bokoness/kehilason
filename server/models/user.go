package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    	string 	  `json:"email"`
	Password    string    `json:"password"`
	FullName    string    `json:"fullName" gorm:"column: full_name"`
	IsSuperuser bool      `json:"isSuperuser" gorm:"is_superuser;default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime: true"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime: true"`
	

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
