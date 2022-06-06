package ping

import "github.com/gin-gonic/gin"

// Ping returns a pong
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
