package db

import (
	"database/sql"
	"log"
)

func GetOrCreateFolder(name string) (int, error) {
	if name == "" {
		return 0, nil
	}
	var id int
	err := DB.QueryRow(`SELECT id FROM folders WHERE name = $1`, name).Scan(&id)
	if err == sql.ErrNoRows {
		err = DB.QueryRow(`INSERT INTO folders (name) VALUES ($1) RETURNING id`, name).Scan(&id)
		if err != nil {
			log.Printf("Error creating folder: %v", err)
			return 0, err
		}
		return id, nil
	} else if err != nil {
		log.Printf("Error finding folder: %v", err)
		return 0, err
	}
	return id, nil
}

func GetAllFolders() ([]Folder, error) {
	rows, err := DB.Query(`SELECT id, name FROM folders`)
	if err != nil {
		log.Printf("Error retrieving folders: %v", err)
		return nil, err
	}
	defer rows.Close()
	var folders []Folder
	for rows.Next() {
		var f Folder
		err := rows.Scan(&f.ID, &f.Name)
		if err != nil {
			log.Printf("Error scanning folder: %v", err)
			return nil, err
		}
		folders = append(folders, f)
	}
	return folders, nil
}

func AddImage(filepath string, folderID int) (int, error) {
	query := `INSERT INTO images (filepath, folder_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := DB.QueryRow(query, filepath, folderID).Scan(&id)
	if err != nil {
		log.Printf("Error adding image: %v", err)
		return 0, err
	}
	return id, nil
}

func GetImageByID(id int) (*Image, error) {
	query := `SELECT id, filepath, upload_date, modification_date, folder_id FROM images WHERE id = $1`
	image := &Image{}
	err := DB.QueryRow(query, id).Scan(&image.ID, &image.Filepath, &image.UploadDate, &image.ModificationDate, &image.FolderID)
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
	query := `SELECT id, filepath, upload_date, modification_date, folder_id FROM images`
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
		err := rows.Scan(&image.ID, &image.Filepath, &image.UploadDate, &image.ModificationDate, &image.FolderID)
		if err != nil {
			log.Printf("Error scanning image: %v", err)
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func GetFolderByID(id int) (*Folder, error) {
	query := `SELECT id, name FROM folders WHERE id = $1`
	folder := &Folder{}
	err := DB.QueryRow(query, id).Scan(&folder.ID, &folder.Name)
	if err != nil {
		log.Printf("Error retrieving folder: %v", err)
		return nil, err
	}
	return folder, nil
}
