package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/api/metadata/:imageId", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "read metadata", "id": c.Param("imageId")})
	})
	r.POST("/api/metadata/:imageId", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "write metadata", "id": c.Param("imageId")})
	})
	return r
}
