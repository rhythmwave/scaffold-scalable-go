// internal/adaptors/validation/file_size_validator.go
package validation

import (
	"chat-backend-general/internal/domain"
)

// FileSizeValidator is a concrete implementation of the FileValidator interface.
type FileSizeValidator struct {
	MaxSize int64 // Maximum allowed file size in bytes
}

// Validate checks if the file size is within the allowed limit.
func (v *FileSizeValidator) Validate(file domain.UploadedFile) error {
	if file.Size > v.MaxSize {
		return domain.ErrFileSizeExceeded
	}
	return nil
}

// NewFileSizeValidator creates a new FileSizeValidator.
func NewFileSizeValidator(maxSize int64) *FileSizeValidator {
	return &FileSizeValidator{
		MaxSize: maxSize,
	}
}
