package router

import (
	"backend/db"
	"github.com/barasher/go-exiftool"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := r.Group("")
	api.Use(assetExistsMiddleware())
	{
		api.GET("/api/metadata/:imageId", func(c *gin.Context) {
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
		api.POST("/api/metadata/:imageId", func(c *gin.Context) {
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
		api.GET("/assets/:imageId", func(c *gin.Context) {
			id := c.Param("imageId")
			file, _ := getFileById(id)
			c.File(file)
		})
		api.DELETE("/api/metadata/:imageId", func(c *gin.Context) {
			id := c.Param("imageId")
			intID, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(400, gin.H{"message": "invalid image ID"})
				return
			}
			err = db.DeleteImageByID(intID)
			if err != nil {
				c.JSON(500, gin.H{"message": "internal server error"})
				return
			}
			c.JSON(200, gin.H{"success": true})
		})
	}
	r.GET("/api/metadata", func(c *gin.Context) {
		images, err := db.GetAllImages()
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
			return
		}
		for i := range images {
			images[i].Filepath = ""
		}
		c.JSON(200, images)
	})

	r.POST("/assets", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": "invalid file"})
			return
		}
		id, err := db.AddImage(file.Filename)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		err = c.SaveUploadedFile(file, filepath.Join("assets", file.Filename))
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		c.JSON(200, gin.H{"id": id})
	})

	return r
}

func assetExistsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("imageId")
		file, err := getFileById(id)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
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
	intID, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}
	image, err := db.GetImageByID(intID)
	if err != nil {
		return "", err
	}
	if image == nil || image.Filepath == "" {
		return "", os.ErrNotExist
	}
	path := filepath.Join("assets", image.Filepath)

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}
