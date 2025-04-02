package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/haseeb492/portfolio-cms/models"
)

func GenerateOTP(length int)  string{
	otp := ""

	for range length {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			fmt.Println("Error generating OTP: ", err)
		}

		otp += n.String()
	}

	return otp
}

func SendOTPEmail(email, otp string)  {
	log.Printf("Sending OTP to %s email: %s", otp, email)
}

func GenerateToken( user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id" : user.ID,
		"email" : user.Email,
		"role" : user.Role,
		"exp" : time.Now().Add(72 * time.Hour).Unix(),
	}
	jwtSecretKey := os.Getenv("JWT_SECRET")
	log.Println("JWT secret: ", jwtSecretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(jwtSecretKey)
	return token.SignedString(secret)
}

func ValidateJWTToken(tokenString string) (jwt.MapClaims, error)  {

	jwtSecret := os.Getenv("JWT_SECRET")
	log.Println("JWT secret: ", jwtSecret)
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is not set")
	}	

	token,err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok:= t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}