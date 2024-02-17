package storage

import (
	"context"
	"mime/multipart"
)

type File struct {
	URL string `json:"url"`
}

type StorageRepo interface {
	UploadFile(ctx context.Context, file multipart.FileHeader) (string, error)
}
