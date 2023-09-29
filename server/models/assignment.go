package models

import (
	"gorm.io/gorm"
	"time"
)

type Assignment struct {
	gorm.Model
	CommunityID string `json:"communityId" gorm:"column:community_id"`
	UserID      uint   `json:"user_id"`
	Start       time.Time
	End         time.Time
	StartHour   time.Time `json:"start_hour"`
	EndHour     time.Time `json:"end_hour"`
}
