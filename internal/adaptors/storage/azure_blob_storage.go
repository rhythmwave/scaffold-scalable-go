package storage

import (
	"chat-backend-general/config"
	"chat-backend-general/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"go.uber.org/zap"
)

// BlobStorageAdapter is an adapter for Azure Blob Storage.
type BlobStorageAdapter struct {
	containerName string
	blobService   *azblob.Client
	logger        *zap.Logger
}

// NewBlobStorageAdapter initializes a new BlobStorageAdapter.
func NewBlobStorageAdapter(cfg *config.Config, logger *zap.Logger) (*BlobStorageAdapter, error) {
	if cfg == nil {
		return nil, errors.New("config is nil")
	}

	// Initialize Azure Blob Storage client
	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", cfg.Storage.Config.AccountName)
	cred, err := azblob.NewSharedKeyCredential(cfg.Storage.Config.AccountName, cfg.Storage.Config.ApiKey)
	if err != nil {
		logger.Error("Failed to create shared key credential", zap.Error(err))
		return nil, fmt.Errorf("failed to create shared key credential: %w", err)
	}

	blobService, err := azblob.NewClientWithSharedKeyCredential(accountURL, cred, nil)
	if err != nil {
		logger.Error("Failed to create Azure Blob Storage client", zap.Error(err))
		return nil, fmt.Errorf("failed to create Azure Blob Storage client: %w", err)
	}

	return &BlobStorageAdapter{
		containerName: cfg.Storage.Config.BucketName,
		blobService:   blobService,
		logger:        logger,
	}, nil
}

// UploadFile uploads a file to Azure Blob Storage.
func (b *BlobStorageAdapter) UploadFile(ctx context.Context, file domain.UploadedFile) error {
	if b.blobService == nil {
		b.logger.Error("blobService is nil")
		return errors.New("blobService is nil")
	}
	if file.File == nil {
		b.logger.Error("file is nil")
		return errors.New("file is nil")
	}

	uploadOpt := &azblob.UploadStreamOptions{}
	_, err := b.blobService.UploadStream(ctx, b.containerName, file.Path, file.File, uploadOpt)
	if err != nil {
		b.logger.Error("Error uploading file", zap.Error(err), zap.String("path", file.Path))
		return fmt.Errorf("failed to upload file: %w", err)
	}

	b.logger.Info("File uploaded successfully", zap.String("path", file.Path))
	return nil
}
