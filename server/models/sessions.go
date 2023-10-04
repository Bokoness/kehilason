package models

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	UserID       uint
	User         User      `gorm:"foreignKey:UserID"`
	RefreshToken string    `json:"refreshToken" gorm:"column:refresh_token"`
	ClientIp     string    `json:"clientIp" gorm:"column:client_ip"`
	IsBlocked    bool      `json:"isBlocked" gorm:"column:is_blocked"`
	ExpiresAt    time.Time `json:"expiresAt" gorm:"column:expires_at"`
}
