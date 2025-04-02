package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/haseeb492/portfolio-cms/controllers"
	"github.com/haseeb492/portfolio-cms/middlewares"
)

func AuthRoutes(router *gin.Engine)  {
	authGroup := router.Group("/auth")

	{
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/submit-otp", controllers.SubmitOtp)
	}

	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		adminGroup.POST("/add-user", controllers.AddUser)
	}
} 