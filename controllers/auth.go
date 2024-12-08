package controllers

import (
	"auth-system/config"
	"auth-system/models"
	"auth-system/utils"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func isValidPhoneNumber(phoneNumber string) bool {
	// Regular expression to match a 10-digit phone number
	regex := regexp.MustCompile(`^\d{10}$`)
	return regex.MatchString(phoneNumber)
}

var limiter = rate.NewLimiter(1, 3) // 1 request per second, burst of 3

func Register(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Please try again later."})
		return
	}
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the phone number
	if !isValidPhoneNumber(user.MobileNumber) {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid phone number. Must be 10 digits."})
		return
	}

	// Check if the user already exists
	usersCollection := config.GetCollection("users")
	_, err := models.FindUserByMobile(usersCollection, user.MobileNumber)
	if err == nil {
		utils.RespondJSON(c, http.StatusConflict, gin.H{"error": "User already registered with this number"})
		return
	}

	// Register the new user
	user.CreatedAt = time.Now()
	if err := models.CreateUser(usersCollection, user); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	mobile := c.PostForm("mobileNumber")

	// Check if the user exists
	usersCollection := config.GetCollection("users")
	_, err := models.FindUserByMobile(usersCollection, mobile)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "User not registered"})
		return
	}

	// Generate OTP
	otp := utils.GenerateOTP()
	otpCollection := config.GetCollection("otps")
	err = models.CreateOTP(otpCollection, models.OTP{
		MobileNumber: mobile,
		OTP:          otp,
		ExpiresAt:    time.Now().Add(5 * time.Minute),
	})
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "OTP sent", "otp": otp}) // Expose OTP only for testing!
}

func ValidateOTP(c *gin.Context) {
	// Extract mobile number and OTP from the request body
	mobileNumber := c.PostForm("mobileNumber")
	otp := c.PostForm("otp")

	// Get the OTP collection from MongoDB
	otpCollection := config.GetCollection("otps")

	// Validate the OTP using the model function
	isValid, err := models.ValidateOTP(otpCollection, mobileNumber, otp)
	if err != nil || !isValid {
		utils.RespondJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "OTP validated successfully"})
}
