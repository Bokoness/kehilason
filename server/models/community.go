package models

import (
	"gorm.io/gorm"
	"time"
)

type Community struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `json:"name"`

	Users []*User `gorm:"many2many:communities_users;"`
}

type CommunitiesUsers struct {
	CommunityID string    `gorm:"primaryKey"`
	UserID      uint      `gorm:"primaryKey"`
	Role        int       `gorm:"default: 2"`
	CreatedAt   time.Time `gorm:"autoCreateTime: true"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime: true"`
}
