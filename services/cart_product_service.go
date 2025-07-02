package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CartProductService interface {
	AddCartProduct()
	GetCartProductsByCart()
	UpdateCartProductById()
	DeleteCartProductById()
}

type CartProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CartProductServiceImpl) AddCartProduct() {

}

func (service *CartProductServiceImpl) GetCartProductsByCart() {

}

func (service *CartProductServiceImpl) UpdateCartProductById() {

}

func (service *CartProductServiceImpl) DeleteCartProductById() {

}

func NewCartProductService(db *gorm.DB, redis *redis.Client) CartProductService {
	return &CartProductServiceImpl{DB: db, Redis: redis}
}
