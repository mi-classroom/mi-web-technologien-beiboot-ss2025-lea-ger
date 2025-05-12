package db

import (
	"database/sql"
	"log"
)

func AddImage(filepath string) (int, error) {
	query := `INSERT INTO images (filepath) VALUES ($1) RETURNING id`
	var id int
	err := DB.QueryRow(query, filepath).Scan(&id)
	if err != nil {
		log.Printf("Error adding image: %v", err)
		return 0, err
	}
	return id, nil
}

func GetImageByID(id int) (*Image, error) {
	query := `SELECT id, filepath, upload_date, modification_date FROM images WHERE id = $1`
	image := &Image{}
	err := DB.QueryRow(query, id).Scan(&image.ID, &image.Filepath, &image.UploadDate, &image.ModificationDate)
	if err != nil {
		log.Printf("Error retrieving image: %v", err)
		return nil, err
	}
	return image, nil
}

func DeleteImageByID(id int) error {
	query := `DELETE FROM images WHERE id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting image: %v", err)
		return err
	}
	return nil
}

func GetAllImages() ([]Image, error) {
	query := `SELECT id, filepath, upload_date, modification_date FROM images`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Error retrieving images: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	var images []Image
	for rows.Next() {
		image := Image{}
		err := rows.Scan(&image.ID, &image.Filepath, &image.UploadDate, &image.ModificationDate)
		if err != nil {
			log.Printf("Error scanning image: %v", err)
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}
