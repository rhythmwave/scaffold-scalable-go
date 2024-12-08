package domain

import "errors"

// FileValidator is an interface for file validation.
type FileValidator interface {
	Validate(file UploadedFile) error
}

// ErrFileTypeInvalid is returned when the file type is invalid.
var ErrFileTypeInvalid = errors.New("invalid file type")

// ErrFileSizeExceeded is returned when the file size exceeds the allowed limit.
var ErrFileSizeExceeded = errors.New("file size exceeded")
