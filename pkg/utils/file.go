package utils

import (
	"context"
	"fmt"
	"io"
	"konsultanku-v2/pkg/databases"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

const (
	storageBucket = "https://firebasestorage.googleapis.com/v0/b/konsultanku-8213e.appspot.com"
)

func GenerateRandomFilename() string {
	shortUUID := uuid.New().String()[:8]
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s", timestamp, shortUUID)
}

func UploadFile(ctx context.Context, file multipart.FileHeader) (string, error) {
	// Determine filename
	filename := GenerateRandomFilename()
	extension := filepath.Ext(file.Filename)
	objectName := fmt.Sprintf("picture/%s%s", filename, extension)

	// Convert file into multipart.File
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// storage.BucketHandler
	bucket, err := databases.StorageFile.Bucket("konsultanku-8213e.appspot.com")
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
