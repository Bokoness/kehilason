package models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	StartHour time.Time `json:"start_hour"`
	EndHour   time.Time `json:"end_hour"`
	gorm.Model
	CommunityID string `json:"communityId" gorm:"column:community_id"`
	UserID      uint   `json:"user_id"`
}
