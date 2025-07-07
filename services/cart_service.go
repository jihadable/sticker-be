package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CartService interface {
	AddCart(cart *models.Cart) (*models.Cart, error)
	GetCartByCustomer(customer_id string) (*models.Cart, error)
}

type CartServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CartServiceImpl) AddCart(cart *models.Cart) (*models.Cart, error) {
	err := service.DB.Create(cart).Error
	if err != nil {
		return nil, err
	}

	return service.GetCartByCustomer(cart.CustomerId)
}

func (service *CartServiceImpl) GetCartByCustomer(customer_id string) (*models.Cart, error) {
	ctx := context.Background()
	cacheKey := "cart:customer:" + customer_id

	cartInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && cartInRedis != "" {
		cart := models.Cart{}
		err = json.Unmarshal([]byte(cartInRedis), &cart)
		if err != nil {
			return nil, err
		}

		return &cart, nil
	}

	cart := models.Cart{}
	err = service.DB.Where("customer_id = ?", customer_id).Preload("Customer").Preload("Products").First(&cart).Error
	if err != nil {
		return nil, err
	}

	cartJSON, err := json.Marshal(cart)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, cartJSON, 30*time.Minute)

	return &cart, nil
}

func NewCartService(db *gorm.DB, redis *redis.Client) CartService {
	return &CartServiceImpl{DB: db, Redis: redis}
}
