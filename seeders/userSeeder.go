package seeders

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/haseeb492/portfolio-cms/models"
)
func SeedAdminUser(db *gorm.DB)  error {
	var count int64
	
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("User already exists. Skipping seeding...")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Haseeb770"), bcrypt.DefaultCost)
	if err != nil {
		return err
	} 

	admin := models.User{
		Name: "Haseeb Ur Rehman",
		Email: "uhaseeb348@gmail.com",
		PasswordHash: string(hashedPassword),
		Role : models.RoleAdmin,
		ContactInfo: "+923224495773",
		IsFirstTime: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err:= db.Create(&admin).Error; err != nil {
		return err
	}

	log.Println("Admin User seeded Successfully")
	return nil
}