package models

import (
	"time"
)

type OTPType string

const (
	OTPTypeLogin         OTPType = "login"
	OTPTypeFirstTime     OTPType = "first_time"
	OTPTypeForgotPassword OTPType = "forgot_password"
)

type UserOTP struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	OTPCode   string    `gorm:"size:100;not null"`
	OTPType   OTPType   `gorm:"type:VARCHAR(50);not null;check:otp_type IN ('login','first_time','forgot_password')"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time
}
