package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductService interface {
	AddProduct()
	GetProductById()
	GetProductsByCategory()
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

func (service *ProductServiceImpl) GetProductsByCategory() {

}

func (service *ProductServiceImpl) GetProducts() {

}

func (service *ProductServiceImpl) UpdateProductById() {

}

func (service *ProductServiceImpl) DeleteProductById() {

}

func NewProductService(db *gorm.DB, redis *redis.Client) ProductService {
	return &ProductServiceImpl{DB: db, Redis: redis}
}
