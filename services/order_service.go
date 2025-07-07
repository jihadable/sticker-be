package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderService interface {
	AddOrder(order *models.Order, orderProducts []*models.OrderProduct) (*models.Order, error)
	GetOrderById(id string) (*models.Order, error)
	GetOrdersByCustomer(customer_id string) ([]*models.Order, error)
	UpdateOrderById(id string, updatedOrder *models.Order) (*models.Order, error)
}

type OrderServiceImpl struct {
	DB                  *gorm.DB
	Redis               *redis.Client
	OrderProductService OrderProductService
}

func (service *OrderServiceImpl) AddOrder(order *models.Order, orderProducts []*models.OrderProduct) (*models.Order, error) {
	err := service.DB.Create(order).Error
	if err != nil {
		return nil, err
	}

	_, err = service.OrderProductService.AddOrderProducts(order.Id, orderProducts)
	if err != nil {
		return nil, err
	}

	order, err = service.GetOrderById(order.Id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (service *OrderServiceImpl) GetOrderById(id string) (*models.Order, error) {
	order := models.Order{}
	err := service.DB.Where("id = ?", id).Preload("Customer").Preload("Products").First(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (service *OrderServiceImpl) GetOrdersByCustomer(customer_id string) ([]*models.Order, error) {
	orders := []*models.Order{}
	err := service.DB.Where("customer_id = ?", customer_id).Preload("Customer").Preload("Products").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (service *OrderServiceImpl) UpdateOrderById(id string, updatedOrder *models.Order) (*models.Order, error) {
	order, err := service.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	order.Status = updatedOrder.Status
	order.TotalPrice = updatedOrder.TotalPrice

	err = service.DB.Save(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func NewOrderService(db *gorm.DB, redis *redis.Client) OrderService {
	return &OrderServiceImpl{DB: db, Redis: redis}
}
