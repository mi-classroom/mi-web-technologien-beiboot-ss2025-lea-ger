package apperrors

import "errors"

var (
	ErrInvalidID      = errors.New("invalid ID")
	ErrImageNotFound  = errors.New("image not found")
	ErrFileNotFound   = errors.New("file not found on disk")
	ErrFolderNotFound = errors.New("folder not found")
)
