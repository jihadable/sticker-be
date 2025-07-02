package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderService interface {
	AddOrder()
	GetOrderById()
	GetOrdersByUser()
	UpdateOrderById()
}

type OrderServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *OrderServiceImpl) AddOrder() {

}

func (service *OrderServiceImpl) GetOrderById() {

}

func (service *OrderServiceImpl) GetOrdersByUser() {

}

func (service *OrderServiceImpl) UpdateOrderById() {

}

func NewOrderService(db *gorm.DB, redis *redis.Client) OrderService {
	return &OrderServiceImpl{DB: db, Redis: redis}
}
