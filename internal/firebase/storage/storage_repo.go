package storage

import (
	"context"
	"fmt"
	"io"
	"konsultanku-v2/pkg/utils"
	"mime/multipart"
	"net/url"
	"path/filepath"

	"firebase.google.com/go/storage"
)

const (
	storageBucket = "https://firebasestorage.googleapis.com/v0/b/konsultanku-8213e.appspot.com"
)

type Repo struct {
	Storage *storage.Client
}

func NewRepo(s *storage.Client) StorageRepo {
	return &Repo{
		Storage: s,
	}
}

func (r *Repo) UploadFile(ctx context.Context, file multipart.FileHeader) (string, error) {

	// Determine filename
	filename := utils.GenerateRandomFilename()
	extension := filepath.Ext(file.Filename)
	objectName := fmt.Sprintf("picture/%s%s", filename, extension)

	// Convert file into multipart.File
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// storage.BucketHandler
	bucket, err := r.Storage.Bucket("konsultanku-8213e.appspot.com")
	if err != nil {
		return "", err
	}

	// Make storage.Writer
	object := bucket.Object(objectName)
	wc := object.NewWriter(ctx)

	// Copy file into Google Cloud Storage
	if _, err := io.Copy(wc, src); err != nil {
		return "", err
	}

	// Close the Google Cloud Storage object writer
	if err := wc.Close(); err != nil {
		return "", err
	}

	// Make URL string
	// example URL : https://firebasestorage.googleapis.com/v0/b/konsultanku-8213e.appspot.com/o/picture%2F20240117220139_2588a317.jpg?alt=media
	encodedPath := url.PathEscape(objectName)
	storageURL := fmt.Sprintf("%s/o/%s?alt=media", storageBucket, encodedPath)
	return storageURL, nil
}
