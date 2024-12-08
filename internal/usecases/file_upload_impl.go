// internal/usecases/file_upload_impl.go
package usecases

import (
	"chat-backend-general/internal/adaptors/storage"
	"chat-backend-general/internal/domain"
	"context"
)

type fileUploadImpl struct {
	storageAdapter *storage.BlobStorageAdapter
}

func NewFileUploadUseCase(storageAdapter *storage.BlobStorageAdapter) FileUploadUseCase {
	return &fileUploadImpl{storageAdapter: storageAdapter}
}

func (f *fileUploadImpl) HandleFileUpload(ctx context.Context, file domain.UploadedFile) error {

	return f.storageAdapter.UploadFile(ctx, file)
}
