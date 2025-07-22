package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CustomProductService interface {
	AddCustomProduct(customProduct *models.CustomProduct, image graphql.Upload) (*models.CustomProduct, error)
	GetCustomProductById(id string) (*models.CustomProduct, error)
	GetCustomProductsByCustomer(customer_id string) ([]*models.CustomProduct, error)
	UpdateCustomProductById(id string, updatedCustomProduct *models.CustomProduct, image *graphql.Upload) (*models.CustomProduct, error)
	DeleteCustomProductById(id string) error
}

type CustomProductServiceImpl struct {
	DB             *gorm.DB
	Redis          *redis.Client
	StorageService StorageService
}

func (service *CustomProductServiceImpl) AddCustomProduct(customProduct *models.CustomProduct, image graphql.Upload) (*models.CustomProduct, error) {
	imageURL, err := service.StorageService.AddFile(image)
	if err != nil {
		return nil, err
	}

	customProduct.ImageURL = *imageURL

	err = service.DB.Create(customProduct).Error
	if err != nil {
		return nil, err
	}

	return service.GetCustomProductById(customProduct.Id)
}

func (service *CustomProductServiceImpl) GetCustomProductById(id string) (*models.CustomProduct, error) {
	ctx := context.Background()
	cacheKey := "custom_product:" + id

	customProductInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && customProductInRedis != "" {
		customProduct := models.CustomProduct{}
		err = json.Unmarshal([]byte(customProductInRedis), &customProduct)
		if err != nil {
			return nil, err
		}

		return &customProduct, nil
	}

	customProduct := models.CustomProduct{}
	err = service.DB.Where("id = ?", id).Preload("Customer").First(&customProduct).Error
	if err != nil {
		return nil, err
	}

	customProductJSON, err := json.Marshal(customProduct)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, customProductJSON, 30*time.Minute)

	return &customProduct, nil
}

func (service *CustomProductServiceImpl) GetCustomProductsByCustomer(customer_id string) ([]*models.CustomProduct, error) {
	ctx := context.Background()
	cacheKey := "custom_product:customer:" + customer_id

	customProductsInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && customProductsInRedis != "" {
		customProducts := []*models.CustomProduct{}
		err = json.Unmarshal([]byte(customProductsInRedis), &customProducts)
		if err != nil {
			return nil, err
		}

		return customProducts, nil
	}

	customProducts := []*models.CustomProduct{}
	err = service.DB.Where("customer_id = ?", customer_id).Preload("Customer").Find(&customProducts).Error
	if err != nil {
		return nil, err
	}

	customProductsJSON, err := json.Marshal(customProducts)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, customProductsJSON, 30*time.Minute)

	return customProducts, nil
}

func (service *CustomProductServiceImpl) UpdateCustomProductById(id string, updatedCustomProduct *models.CustomProduct, image *graphql.Upload) (*models.CustomProduct, error) {
	customProduct, err := service.GetCustomProductById(id)
	if err != nil {
		return nil, err
	}

	if image != nil {
		err = service.StorageService.DeleteFile(customProduct.ImageURL)
		if err != nil {
			return nil, err
		}

		imageURL, err := service.StorageService.AddFile(*image)
		if err != nil {
			return nil, err
		}

		customProduct.ImageURL = *imageURL
	}

	customProduct.Name = updatedCustomProduct.Name

	err = service.DB.Save(customProduct).Error
	if err != nil {
		return nil, err
	}

	cacheKeys := []string{"custom_product:" + customProduct.Id, "custom_product:customer:" + customProduct.CustomerId}
	service.Redis.Del(context.Background(), cacheKeys...)

	return service.GetCustomProductById(customProduct.Id)
}

func (service *CustomProductServiceImpl) DeleteCustomProductById(id string) error {
	customProduct, err := service.GetCustomProductById(id)
	if err != nil {
		return err
	}

	err = service.DB.Delete(customProduct).Error
	if err != nil {
		return err
	}

	cacheKeys := []string{"custom_product:" + id, "custom_product:customer" + customProduct.CustomerId}
	service.Redis.Del(context.Background(), cacheKeys...)

	return nil
}

func NewCustomProductService(db *gorm.DB, redis *redis.Client) CustomProductService {
	return &CustomProductServiceImpl{DB: db, Redis: redis, StorageService: NewStorageService()}
}
