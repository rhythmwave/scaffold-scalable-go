package domain

import (
	"io"
)

type UploadedFile struct {
	Name        string
	ContentType string
	File        io.Reader
	Size        int64
	Path        string
}
