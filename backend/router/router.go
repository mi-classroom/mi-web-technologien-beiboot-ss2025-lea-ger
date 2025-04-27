package router

import (
	"github.com/barasher/go-exiftool"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"runtime"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := r.Group("/api/metadata/:imageId")
	api.Use(assetExistsMiddleware())
	{
		api.GET("", func(c *gin.Context) {
			file, _ := getFileById(c.Param("imageId"))
			et, err := exiftool.NewExiftool()
			if err != nil {
				c.JSON(500, gin.H{"message": "internal server error: " + err.Error()})
				return
			}
			fileInfos := et.ExtractMetadata(file)
			if len(fileInfos) > 0 {
				c.JSON(200, fileInfos[0].Fields)
				return
			}
			c.JSON(404, gin.H{"message": "metadata not found", "id": c.Param("imageId")})
		})
		api.POST("", func(c *gin.Context) {
			file, _ := getFileById(c.Param("imageId"))
			et, err := exiftool.NewExiftool()
			defer et.Close()
			if err != nil {
				c.JSON(500, gin.H{"message": "internal server error: " + err.Error()})
				return
			}

			var requestBody map[string]interface{}
			if err := c.ShouldBindJSON(&requestBody); err != nil {
				c.JSON(400, gin.H{"message": "invalid request body: " + err.Error()})
				return
			}

			fileInfos := et.ExtractMetadata(file)
			if len(fileInfos) == 0 {
				c.JSON(404, gin.H{"message": "metadata not found", "id": c.Param("imageId")})
				return
			}

			for key, value := range requestBody {
				if key == "FileName" {
					continue
				}
				fileInfos[0].Fields[key] = value
			}

			et.WriteMetadata(fileInfos)

			updatedFileInfos := et.ExtractMetadata(file)
			if updatedFileInfos[0].Err != nil {
				c.JSON(500, gin.H{"message": "internal server error: " + err.Error()})
				return
			}

			c.JSON(200, gin.H{"success": true})
		})
	}
	return r
}

func assetExistsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("imageId")
		file, err := getFileById(id)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			c.Abort()
			return
		}
		if file == "" {
			c.JSON(404, gin.H{"message": "image not found", "id": id})
			c.Abort()
			return
		}
		c.Next()
	}
}

func getFileById(id string) (string, error) {
	_, b, _, _ := runtime.Caller(0) // path to current file
	basePath := filepath.Join(filepath.Dir(b), "../assets")

	dir, err := os.ReadDir(basePath)
	if err != nil {
		return "", err
	}
	for _, file := range dir {
		name := file.Name()
		fileNameWithoutExt := name[:len(name)-len(filepath.Ext(name))]
		if fileNameWithoutExt == id {
			fullPath := filepath.Join(basePath, name)
			if _, err := os.Stat(fullPath); err == nil {
				return fullPath, nil
			}
		}
	}
	return "", nil
}
