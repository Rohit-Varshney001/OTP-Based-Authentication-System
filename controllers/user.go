package controllers

import (
	"auth-system/config"
	"auth-system/models"
	"auth-system/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetUserDetails(c *gin.Context) {
	mobileNumber := c.Param("mobileNumber") // Extract from URL or session

	usersCollection := config.GetCollection("users")
	user, err := models.FindUserByMobile(usersCollection, mobileNumber)

	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{
		"name":              user.Name,
		"mobileNumber":      user.MobileNumber,
		"deviceFingerprint": user.DeviceFingerprint,
		"createdAt":         user.CreatedAt,
	})
}

func UpdateUserDetails(c *gin.Context) {
	mobileNumber := c.Param("mobileNumber") // Extract mobile number from the URL
	var updateData map[string]interface{}   // Data to be updated

	// Validate input
	if mobileNumber == "" {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	// Check if the user exists
	usersCollection := config.GetCollection("users")
	_, err := models.FindUserByMobile(usersCollection, mobileNumber)
	if err != nil {
		utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "User not available"})
		return
	}

	// Bind the update data
	if err := c.BindJSON(&updateData); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update user details
	_, err = usersCollection.UpdateOne(
		context.TODO(),
		bson.M{"mobileNumber": mobileNumber},
		bson.M{"$set": updateData},
	)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	utils.RespondJSON(c, http.StatusOK, gin.H{"message": "User details updated successfully"})
}
