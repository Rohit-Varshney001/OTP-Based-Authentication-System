package routes

import (
	"auth-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/user/:mobileNumber", controllers.GetUserDetails)    // Fetch user details
	router.PUT("/user/:mobileNumber", controllers.UpdateUserDetails) // Update user details
	router.POST("/validate-otp", controllers.ValidateOTP)            // Validate OTP

}
