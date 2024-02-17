package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomFilename() string {
	shortUUID := uuid.New().String()[:8]
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s", timestamp, shortUUID)
}
