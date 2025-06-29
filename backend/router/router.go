package router

import (
	"backend/db"
	apperrors "backend/errors"
	"errors"
	"fmt"
	"github.com/barasher/go-exiftool"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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

			if fileInfos[0].Fields == nil {
				fileInfos[0].Fields = make(map[string]interface{})
			}

			for key, value := range requestBody {
				if key == "FileName" {
					continue
				}
				fileInfos[0].Fields[key] = value
			}

			et.WriteMetadata(fileInfos)

			updatedFileInfos := et.ExtractMetadata(file)
			if len(updatedFileInfos) == 0 || updatedFileInfos[0].Err != nil {
				c.JSON(500, gin.H{"message": "internal server error: " + updatedFileInfos[0].Err.Error()})
				return
			}

			c.JSON(200, gin.H{"success": true})
		})
		api.GET("/assets/:imageId", func(c *gin.Context) {
			id := c.Param("imageId")
			file, _ := getFileById(id)
			c.File(file)
		})
		api.DELETE("/assets/:imageId", func(c *gin.Context) {
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
		_, err := getFileById(id)
		if err != nil {
			var status int
			switch {
			case errors.Is(err, apperrors.ErrInvalidID):
				status = http.StatusBadRequest
			case errors.Is(err, apperrors.ErrImageNotFound),
				errors.Is(err, apperrors.ErrFileNotFound):
				status = http.StatusNotFound
			default:
				status = http.StatusInternalServerError
			}
			c.JSON(status, gin.H{"message": err.Error()})
			return
		}
		c.Next()
	}
}

func getFileById(id string) (string, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("%w: %v", apperrors.ErrInvalidID, err)
	}

	image, err := db.GetImageByID(intID)
	if image == nil || image.Filepath == "" {
		return "", apperrors.ErrImageNotFound
	}

	path := filepath.Join("assets", image.Filepath)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "", apperrors.ErrFileNotFound
		}
		return "", fmt.Errorf("error checking file: %w", err)
	}

	return path, nil
}
