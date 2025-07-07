package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderProductService interface {
	AddOrderProducts(order_id string, orderProducts []*models.OrderProduct) ([]*models.OrderProduct, error)
}

type OrderProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *OrderProductServiceImpl) AddOrderProducts(order_id string, orderProducts []*models.OrderProduct) ([]*models.OrderProduct, error) {
	for _, orderProduct := range orderProducts {
		orderProduct.OrderId = order_id
	}

	err := service.DB.Create(orderProducts).Error
	if err != nil {
		return nil, err
	}

	return orderProducts, nil
}

func NewOrderProductService(db *gorm.DB, redis *redis.Client) OrderProductService {
	return &OrderProductServiceImpl{DB: db, Redis: redis}
}
