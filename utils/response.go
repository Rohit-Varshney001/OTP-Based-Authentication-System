package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, status int, data gin.H) {
	c.JSON(status, data)
}
