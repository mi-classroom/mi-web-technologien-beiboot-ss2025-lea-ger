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
		folderIdStr := c.Query("folder_id")
		if folderIdStr != "" {
			folderId, err := strconv.Atoi(folderIdStr)
			if err != nil {
				c.JSON(400, gin.H{"message": "invalid folder_id"})
				return
			}
			images, err := db.GetImagesByFolderID(folderId)
			if err != nil {
				c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
				return
			}
			c.JSON(200, images)
			return
		}
		images, err := db.GetAllImages()
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
			return
		}
		c.JSON(200, images)
	})

	r.GET("/api/folders", func(c *gin.Context) {
		folders, err := db.GetAllFolders()
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
			return
		}
		c.JSON(200, folders)
	})

	r.POST("/api/folders", func(c *gin.Context) {
		var folder db.Folder
		if err := c.ShouldBindJSON(&folder); err != nil {
			c.JSON(400, gin.H{"message": "invalid request body"})
			return
		}
		if folder.Name == "" {
			c.JSON(400, gin.H{"message": "folder name cannot be empty"})
			return
		}
		id, err := db.GetOrCreateFolder(folder.Name)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"id": id})
	})

	r.POST("/assets", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": "invalid file"})
			return
		}
		folderIDStr := c.PostForm("folder_id")
		folderId, err := strconv.Atoi(folderIDStr)
		if err != nil {
			folderId = 0
		}
		id, err := db.AddImage(file.Filename, folderId)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		targetDir := "assets"
		if folderId != 0 {
			folder, err := db.GetFolderByID(folderId)
			if err != nil {
				c.JSON(500, gin.H{"message": "internal server error"})
				return
			}
			targetDir = filepath.Join("assets", folder.Name)
			err = os.MkdirAll(targetDir, os.ModePerm)
			if err != nil {
				c.JSON(500, gin.H{"message": "could not create folder"})
				return
			}
		}
		targetPath := filepath.Join(targetDir, file.Filename)
		err = c.SaveUploadedFile(file, targetPath)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		response := gin.H{"id": id}
		if folderId != 0 {
			response["folder_id"] = folderId
		}
		c.JSON(200, response)
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

	folderName := ""
	if image.FolderId != 0 {
		folder, err := db.GetFolderByID(image.FolderId)
		if err != nil {
			return "", apperrors.ErrFolderNotFound
		}
		folderName = folder.Name
	}
	var path string
	if folderName != "" {
		path = filepath.Join("assets", folderName, image.Filepath)
	} else {
		path = filepath.Join("assets", image.Filepath)
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "", apperrors.ErrFileNotFound
		}
		return "", fmt.Errorf("error checking file: %w", err)
	}

	return path, nil
}
