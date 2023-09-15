package models

import (
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `json:"name"`
}
