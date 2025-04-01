package models

import (
	"time"
)

type Article struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text;not null"`
	ImageURL    string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
