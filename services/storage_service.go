package services

import storage_go "github.com/supabase-community/storage-go"

type StorageService interface {
	AddFile() (string, error)
	DeleteFile() error
}

type StorageServiceImpl struct {
	Storage *storage_go.Client
}

func (service *StorageServiceImpl) AddFile() (string, error) {
	panic("")
}

func (service *StorageServiceImpl) DeleteFile() error {
	panic("")
}

func NewStorageService() StorageService {
	return &StorageServiceImpl{Storage: storage_go.NewClient("", "", map[string]string{})}
}
