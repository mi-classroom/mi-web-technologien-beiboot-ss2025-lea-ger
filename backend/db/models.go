package db

import "time"

type Image struct {
	ID               int       `json:"id"`
	Filepath         string    `json:"filepath"`
	UploadDate       time.Time `json:"upload_date"`
	ModificationDate time.Time `json:"modification_date"`
}
