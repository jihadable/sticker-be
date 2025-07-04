package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type StorageService interface {
	AddFile(file graphql.Upload) (string, error)
	DeleteFile(fileName string) error
}

type StorageServiceImpl struct {
	Storage *storage_go.Client
	Bucket  string
}

func (service *StorageServiceImpl) AddFile(file graphql.Upload) (string, error) {
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.NewString() + ext

	data, err := service.Storage.UploadFile(service.Bucket, newFileName, file.File, storage_go.FileOptions{ContentType: &file.ContentType})

	fmt.Println("data:", data)
	fmt.Println("error:", err)
	if err != nil {
		return "", err
	}

	return newFileName, nil
}

func (service *StorageServiceImpl) DeleteFile(fileName string) error {
	_, err := service.Storage.RemoveFile(service.Bucket, []string{fileName})

	return err
}

func NewStorageService() StorageService {
	storageURL := os.Getenv("STORAGE_URL")
	storage_api_key := os.Getenv("STORAGE_API_KEY")
	bucket := os.Getenv("STORAGE_BUCKET")

	// fmt.Println(storageURL)
	// fmt.Println(storage_api_key)
	// fmt.Println(bucket)

	return &StorageServiceImpl{
		Storage: storage_go.NewClient(storageURL, storage_api_key, nil),
		Bucket:  bucket,
	}
}
