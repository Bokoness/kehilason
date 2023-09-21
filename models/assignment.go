package models

import (
	"gorm.io/gorm"
	"time"
)

type Assignment struct {
	gorm.Model
	CommunityID string `json:"community_id"`
	UserID      string `json:"user_id"`
	Start       time.Time
	End         time.Time
	StartHour   time.Time `json:"start_hour"`
	EndHour     time.Time `json:"end_hour"`
}
