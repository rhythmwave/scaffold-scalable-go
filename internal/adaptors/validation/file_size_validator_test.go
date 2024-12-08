package validation

import (
	"testing"

	"chat-backend-general/internal/domain"
)

func TestFileSizeValidator_Validate(t *testing.T) {
	validator := &FileSizeValidator{MaxSize: 1024}

	tests := []struct {
		name     string
		file     domain.UploadedFile
		expected error
	}{
		{
			name:     "valid file size",
			file:     domain.UploadedFile{Size: 512},
			expected: nil,
		},
		{
			name:     "exceeded file size",
			file:     domain.UploadedFile{Size: 2048},
			expected: domain.ErrFileSizeExceeded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.file)
			if err != tt.expected {
				t.Errorf("FileSizeValidator.Validate(%v) = %v, want %v", tt.file, err, tt.expected)
			}
		})
	}
}
