package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/haseeb492/portfolio-cms/controllers"
)

func AuthRoutes(router *gin.Engine)  {
	authGroup := router.Group("/auth")

	{
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/submit-otp", controllers.SubmitOtp)
	}
}