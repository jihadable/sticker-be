package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductService interface {
	AddProduct()
	GetProductById(id string) (*models.Product, error)
	GetProductsByCategory(category_id string) ([]*models.Product, error)
	GetProducts() ([]*models.Product, error)
	UpdateProductById()
	DeleteProductById(id string) error
}

type ProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *ProductServiceImpl) AddProduct() {

}

func (service *ProductServiceImpl) GetProductById(id string) (*models.Product, error) {
	ctx := context.Background()
	cacheKey := "product:" + id

	productInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && productInRedis != "" {
		product := models.Product{}
		err := json.Unmarshal([]byte(productInRedis), &product)
		if err != nil {
			return nil, err
		}

		return &product, nil
	}

	product := models.Product{}
	err = service.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(product)
	if err == nil {
		service.Redis.Set(ctx, cacheKey, productJSON, 5*time.Minute)
	}

	return &product, nil
}

func (service *ProductServiceImpl) GetProductsByCategory(category_id string) ([]*models.Product, error) {
	ctx := context.Background()
	cacheKey := "product:category:" + category_id

	productInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && productInRedis != "" {
		products := []*models.Product{}
		err := json.Unmarshal([]byte(productInRedis), &products)
		if err != nil {
			return nil, err
		}

		return products, nil
	}

	products := []*models.Product{}
	err = service.DB.Where("category_id = ?", category_id).First(&products).Error
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(products)
	if err == nil {
		service.Redis.Set(ctx, cacheKey, productJSON, 5*time.Minute)
	}

	return products, nil
}

func (service *ProductServiceImpl) GetProducts() ([]*models.Product, error) {
	products := []*models.Product{}

	err := service.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductServiceImpl) UpdateProductById() {

}

func (service *ProductServiceImpl) DeleteProductById(id string) error {
	product, err := service.GetProductById(id)
	if err != nil {
		return err
	}

	err = service.DB.Delete(&product).Error

	// Hapus dari Redis
	cacheKey := "product:" + id
	service.Redis.Del(context.Background(), cacheKey)

	// (Opsional) Hapus cache by category jika kamu cache per kategori
	// categoryKey := "product:category:" + product.C
	// service.Redis.Del(context.Background(), categoryKey)

	return err
}

func NewProductService(db *gorm.DB, redis *redis.Client) ProductService {
	return &ProductServiceImpl{DB: db, Redis: redis}
}
