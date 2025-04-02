package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleEditor UserRole = "editor"
)

var validRolesSet = map[UserRole]struct{}{
	RoleAdmin:    {},
	RoleEditor:   {},
}

func (r UserRole) IsValid() bool {
	_, ok := validRolesSet[r]
	return ok
}

type User struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `gorm:"size:255;not null"`
	Email        string         `gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string         `gorm:"size:255"`
	Role         UserRole       `gorm:"type:VARCHAR(50);not null;default:'editor';check:role IN ('admin','editor')"`
	ContactInfo  string         `gorm:"size:255"`
	IsFirstTime  bool           `gorm:"default:true"`
}
