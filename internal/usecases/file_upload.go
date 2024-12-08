package usecases

import (
	"context"

	"chat-backend-general/internal/domain"
)

type FileUploadUseCase interface {
	HandleFileUpload(ctx context.Context, file domain.UploadedFile) error
}
