package utils

import "github.com/gin-gonic/gin"


func RespondJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}