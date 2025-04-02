package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haseeb492/portfolio-cms/lib/utils"
	"github.com/haseeb492/portfolio-cms/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"` 
}

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}


func Login(c *gin.Context)  {
	var request LoginRequest
	if err:= c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid request format"})
		return
	}

	dbInstance, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error"})
		log.Fatalf("Error in login contoller: %v", "DB connection not established")
		return
	}
	db := dbInstance.(*gorm.DB)

	var user models.User

	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "User not found"})
		return
	}
	// first time login
	if user.PasswordHash == "" {
		if request.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "OTP is required in the password field for first time login"})
			return
		}

		var otpEntry models.UserOTP
		if err:= db.Where("user_id = ? AND otp_code = ? AND otp_type = ? AND used = false",
		user.ID, request.Password, "first_time").Order("created_at desc").First(&otpEntry).Error; err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid OTP"})
			return
		}

		if time.Now().After(otpEntry.ExpiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "OTP expired"})
			return
		}

		otpEntry.Used = true
		db.Save(&otpEntry)

		//generate jwt token
		token, err := utils.GenerateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			log.Fatalf("Error in login controller: %v", err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user" : user,
			"token" : token,
			"isFirstTimeLogin" : true,
		})
		return
	}
	// normal login flow

	if err:= bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid paswword"})
		return
	}

	otpCode := utils.GenerateOTP(8)
	otpEntry := models.UserOTP{
		UserID: user.ID,
		OTPCode: otpCode,
		OTPType: "login",
		ExpiresAt: time.Now().Add(2 * time.Minute),
		Used: false,
		CreatedAt: time.Now(),
	}

	if err:= db.Create(&otpEntry).Error; err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to create otp"})
		log.Fatalf("Error in login controller: %v", err.Error())
		return
	}

	utils.SendOTPEmail(user.Email, otpCode)
	c.JSON(http.StatusOK, gin.H{
		"message" : "OTP sent to your email",
		"user" : user,
		"isFirstTimeLogin" : false,
	})
}

func SubmitOtp(c *gin.Context)  {
	var request OTPRequest
	if err:= c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	dbInstance, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error"})
		log.Fatalf("Database connection not established")
		return
	}

	db := dbInstance.(*gorm.DB)

	var user models.User
	if err:= db.Where("email = ?", request.Email).First(&user).Error; err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "User not found"})
		return
	}

	var otpEntry models.UserOTP
	if err:= db.Where("user_id = ? AND otp_code = ? AND otp_type = ? AND used = false",
		user.ID, request.OTP, "login").Order("created_at desc").First(&otpEntry).Error; err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid otp"})
		return
	}

	otpEntry.Used = true
	db.Save(&otpEntry)

	//generate jwt token

	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error"})
		log.Fatalf("Error in submit otp controller: %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token" : token,
		"user" : user,
		"isFirstTimeLogin" : false,
	})
}

func SetPassword(c *gin.Context)  {
	
}