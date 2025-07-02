package services

import (
	"github.com/jihadable/sticker-be/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductService interface {
	AddProduct()
	GetProductById()
	GetProducts()
	UpdateProductById()
	DeleteProductById()
}

type ProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *ProductServiceImpl) AddProduct() {

}

func (service *ProductServiceImpl) GetProductById() {

}

func (service *ProductServiceImpl) GetProducts() {

}

func (service *ProductServiceImpl) UpdateProductById() {

}

func (service *ProductServiceImpl) DeleteProductById() {

}

func NewProductService() ProductService {
	return &ProductServiceImpl{DB: config.DB(), Redis: config.Redis()}
}
