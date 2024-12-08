package validation

import (
	"testing"

	"chat-backend-general/internal/domain"
)

func TestFileTypeValidator_Validate(t *testing.T) {
	validator := &FileTypeValidator{AllowedTypes: []string{"image/jpeg", "image/png"}}

	tests := []struct {
		name     string
		file     domain.UploadedFile
		expected error
	}{
		{
			name:     "valid file type",
			file:     domain.UploadedFile{ContentType: "image/jpeg"},
			expected: nil,
		},
		{
			name:     "invalid file type",
			file:     domain.UploadedFile{ContentType: "application/pdf"},
			expected: domain.ErrFileTypeInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.file)
			if err != tt.expected {
				t.Errorf("FileTypeValidator.Validate(%v) = %v, want %v", tt.file, err, tt.expected)
			}
		})
	}
}
