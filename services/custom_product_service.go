package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CustomProductService interface {
	AddCustomProduct()
	GetCustomProductById()
	GetCustomProductsByUser()
	UpdateCustomProductById()
	DeleteCustomProductById()
}

type CustomProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CustomProductServiceImpl) AddCustomProduct() {

}

func (service *CustomProductServiceImpl) GetCustomProductById() {

}

func (service *CustomProductServiceImpl) GetCustomProductsByUser() {

}

func (service *CustomProductServiceImpl) UpdateCustomProductById() {

}

func (service *CustomProductServiceImpl) DeleteCustomProductById() {

}

func NewCustomProductService(db *gorm.DB, redis *redis.Client) CustomProductService {
	return &CustomProductServiceImpl{DB: db, Redis: redis}
}
