package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CartService interface {
	AddCart()
	GetCartByUser()
}

type CartServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CartServiceImpl) AddCart() {

}

func (service *CartServiceImpl) GetCartByUser() {

}

func NewCartService(db *gorm.DB, redis *redis.Client) CartService {
	return &CartServiceImpl{DB: db, Redis: redis}
}
