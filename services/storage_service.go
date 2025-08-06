package services

import (
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type StorageService interface {
	AddFile(file graphql.Upload) (*string, error)
	DeleteFile(fileName string) error
	GetPublicFileURL(filePath string) (string, error)
}

type StorageServiceImpl struct {
	Storage *storage_go.Client
	Bucket  string
}

func (service *StorageServiceImpl) AddFile(file graphql.Upload) (*string, error) {
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.NewString() + ext
	contentType := getContentType(ext)

	_, err := service.Storage.UploadFile(service.Bucket, newFileName, file.File, storage_go.FileOptions{ContentType: &contentType})

	if err != nil {
		return nil, err
	}

	return &newFileName, nil
}

func (service *StorageServiceImpl) DeleteFile(fileName string) error {
	_, err := service.Storage.RemoveFile(service.Bucket, []string{fileName})

	return err
}

func (service *StorageServiceImpl) GetPublicFileURL(filePath string) (string, error) {
	publicURL := service.Storage.GetPublicUrl(service.Bucket, filePath)

	return publicURL.SignedURL, nil
}

func getContentType(ext string) string {
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}

func NewStorageService() StorageService {
	storageURL := os.Getenv("STORAGE_URL")
	storage_api_key := os.Getenv("STORAGE_API_KEY")
	bucket := os.Getenv("STORAGE_BUCKET")

	return &StorageServiceImpl{
		Storage: storage_go.NewClient(storageURL, storage_api_key, nil),
		Bucket:  bucket,
	}
}
