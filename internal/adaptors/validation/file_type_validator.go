// internal/adaptors/validation/file_type_validator.go
package validation

import (
	"chat-backend-general/internal/domain"
)

// FileTypeValidator is a concrete implementation of the FileValidator interface.
type FileTypeValidator struct {
	AllowedTypes []string // List of allowed file types
}

// Validate checks if the file type is one of the allowed types.
func (v *FileTypeValidator) Validate(file domain.UploadedFile) error {
	for _, allowedType := range v.AllowedTypes {
		if file.ContentType == allowedType {
			return nil
		}
	}
	return domain.ErrFileTypeInvalid
}

// NewFileTypeValidator creates a new FileTypeValidator.
func NewFileTypeValidator(allowedTypes []string) *FileTypeValidator {
	return &FileTypeValidator{
		AllowedTypes: allowedTypes,
	}
}
