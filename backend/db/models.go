package db

import "time"

type Folder struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID               int       `json:"id"`
	Filepath         string    `json:"filepath"`
	UploadDate       time.Time `json:"upload_date"`
	ModificationDate time.Time `json:"modification_date"`
	FolderID         int       `json:"folder_id"`
}
