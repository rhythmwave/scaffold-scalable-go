package domain

import "context"

// FileUploadUseCase defines the interface for uploading files.
type FileUploadUseCase interface {
	HandleFileUpload(ctx context.Context, file UploadedFile) error
}
