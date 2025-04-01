package models

import (
	"time"

	"github.com/lib/pq"
)

type Project struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"not null"`
	Title       string         `gorm:"size:255;not null"`
	Description string         `gorm:"type:text;not null"`
	TechStack   pq.StringArray `gorm:"type:text[]"`
	Images      pq.StringArray `gorm:"type:text[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
