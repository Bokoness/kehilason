package models

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	Users []*User `gorm:"many2many:communities_users;"`
}

type CommunitiesUsers struct {
	CreatedAt   time.Time `gorm:"autoCreateTime: true"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime: true"`
	CommunityID string    `json:"communityId" gorm:"primaryKey"`
	UserID      uint      `json:"userId" gorm:"primaryKey"`
	Role        int       `json:"role" gorm:"default: 2"`
}
