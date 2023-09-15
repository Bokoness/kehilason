package models

import (
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `json:"name"`

	Users []*User `gorm:"many2many:campaigns_users;"`
}

type CampaignsUsers struct {
	CampaignID int       `gorm:"primaryKey"`
	UserID     string    `gorm:"primaryKey"`
	Role       int       `gorm:"default: 2"`
	CreatedAt  time.Time `gorm:"autoCreateTime: true"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime: true"`
}
