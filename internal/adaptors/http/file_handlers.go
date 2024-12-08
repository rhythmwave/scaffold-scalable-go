// internal/adaptors/http/file_handlers.go
package http

import (
	"chat-backend-general/internal/domain"
	"chat-backend-general/internal/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileUploadUseCase usecases.FileUploadUseCase
	fileValidators    []domain.FileValidator
}

func NewFileHandler(fileUploadUseCase usecases.FileUploadUseCase, fileValidators []domain.FileValidator) *FileHandler {
	return &FileHandler{fileUploadUseCase: fileUploadUseCase, fileValidators: fileValidators}
}

func (f *FileHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	uname := c.PostForm("username")
	chatid := c.PostForm("chatid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	// Read file data
	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	defer fileData.Close()

	contPath := uname + "/" + chatid + "/" + file.Filename
	// Create domain.UploadedFile instance
	uploadedFile := domain.UploadedFile{
		Name:        file.Filename,
		ContentType: file.Header.Get("Content-Type"),
		File:        fileData,
		Size:        file.Size,
		Path:        contPath,
	}

	// Validate file
	for _, validator := range f.fileValidators {
		if err := validator.Validate(uploadedFile); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Handle file upload using the use case
	if err := f.fileUploadUseCase.HandleFileUpload(context.Background(), uploadedFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
