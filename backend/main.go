package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/metadata/:imageId", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "read metadata", "id": c.Param("imageId")})
	})

	r.POST("/api/metadata/:imageId", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "write metadata", "id": c.Param("imageId")})
	})

	r.Run() // Default on :8080
}
