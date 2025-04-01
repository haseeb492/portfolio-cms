package models

import "time"

type Contact struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;not null"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}
