package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/haseeb492/portfolio-cms/models"
	"github.com/haseeb492/portfolio-cms/seeders"
)

func main () {
	dbConnectionString := os.Getenv("DATABASE_URL")
	if dbConnectionString == "" {
		log.Fatal("Database url in env not set")
	}

	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.UserOTP{},
		&models.Project{},
		&models.Article{},
		&models.Contact{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
 
	if err:= seeders.SeedAdminUser(db); err != nil {
		log.Fatalf("Failed to seed admin User: %v", err) 
	}

	router := gin.Default()
 
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message" : "Hello from portfolio CMS"})
	}) 

	log.Println("server is running on port: 808")	
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	
}