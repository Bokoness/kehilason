package models

import "time"

type User struct {
	Email     string    `json:"email" gorm:"primaryKey"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName" gorm:"column: full_name"`
	CreatedAt time.Time `gorm:"autoCreateTime: true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime: true"`

	Campaigns []*Campaign `gorm:"many2many:campaigns_users"`
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
