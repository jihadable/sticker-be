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

type ProductService interface {
	AddProduct(product *models.Product, image graphql.Upload) (*models.Product, error)
	GetProductById(id string) (*models.Product, error)
	GetProductsByCategory(category_id string) ([]*models.Product, error)
	GetProducts() ([]*models.Product, error)
	UpdateProductById(id string, updatedProduct *models.Product, image *graphql.Upload) (*models.Product, error)
	DeleteProductById(id string) error
}

type ProductServiceImpl struct {
	DB             *gorm.DB
	Redis          *redis.Client
	StorageService StorageService
}

func (service *ProductServiceImpl) AddProduct(product *models.Product, image graphql.Upload) (*models.Product, error) {
	imageURL, err := service.StorageService.AddFile(image)
	if err != nil {
		return nil, err
	}

	product.ImageURL = *imageURL

	err = service.DB.Save(product).Error
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(context.Background(), "product:"+product.Id, productJSON, 5*time.Minute)

	return product, nil
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
	err = service.DB.Where("id = ?", id).Preload("Categories").First(&product).Error
	if err != nil {
		return nil, err
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, productJSON, 5*time.Minute)

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
	err = service.DB.Where("category_id = ?", category_id).Preload("Categories").First(&products).Error
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

	err := service.DB.Preload("Categories").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *ProductServiceImpl) UpdateProductById(id string, updatedProduct *models.Product, image *graphql.Upload) (*models.Product, error) {
	product, err := service.GetProductById(id)
	if err != nil {
		return nil, err
	}

	if image != nil {
		err = service.StorageService.DeleteFile(product.ImageURL)
		if err != nil {
			return nil, err
		}

		imageURL, err := service.StorageService.AddFile(*image)
		if err != nil {
			return nil, err
		}
		product.ImageURL = *imageURL
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Stock = updatedProduct.Stock
	product.Description = updatedProduct.Description

	err = service.DB.Save(&product).Error
	if err != nil {
		return nil, err
	}

	cacheKeys := []string{"product:" + id}

	if product.Categories != nil {
		for _, category := range product.Categories {
			cacheKeys = append(cacheKeys, "product:category:"+category.Id)
		}
	}
	service.Redis.Del(context.Background(), cacheKeys...)

	return product, nil
}

func (service *ProductServiceImpl) DeleteProductById(id string) error {
	product, err := service.GetProductById(id)
	if err != nil {
		return err
	}

	err = service.DB.Delete(&product).Error

	cacheKeys := []string{"product:" + id}

	if product.Categories != nil {
		for _, category := range product.Categories {
			cacheKeys = append(cacheKeys, "product:category:"+category.Id)
		}
	}
	service.Redis.Del(context.Background(), cacheKeys...)

	return err
}

func NewProductService(db *gorm.DB, redis *redis.Client) ProductService {
	return &ProductServiceImpl{DB: db, Redis: redis, StorageService: NewStorageService()}
}
