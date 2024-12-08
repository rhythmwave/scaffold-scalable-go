package domain

type FileRepository interface {
	SaveFile(file UploadedFile) error
	GetFile(id string) (UploadedFile, error)
}
