package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CategoryService interface {
	AddCategory(category *models.Category) (*models.Category, error)
	GetCategoryById(id string) (*models.Category, error)
	GetCategories() ([]*models.Category, error)
	DeleteCategoryById(id string) error
}

type CategoryServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CategoryServiceImpl) AddCategory(category *models.Category) (*models.Category, error) {
	err := service.DB.Create(category).Error
	if err != nil {
		return nil, err
	}

	return service.GetCategoryById(category.Id)
}

func (service *CategoryServiceImpl) GetCategoryById(id string) (*models.Category, error) {
	ctx := context.Background()
	cacheKey := "category:" + id

	categoryIdRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && categoryIdRedis != "" {
		category := models.Category{}
		err := json.Unmarshal([]byte(categoryIdRedis), &category)
		if err != nil {
			return nil, err
		}

		return &category, nil
	}

	category := models.Category{}
	err = service.DB.Where("id = ?", id).Preload("Products").First(&category).Error
	if err != nil {
		return nil, err
	}

	categoryJSON, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, categoryJSON, 24*time.Hour)

	return &category, nil
}

func (service *CategoryServiceImpl) GetCategories() ([]*models.Category, error) {
	ctx := context.Background()
	cacheKey := "categories"
	categoriesInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && categoriesInRedis != "" {
		categories := []*models.Category{}
		err = json.Unmarshal([]byte(categoriesInRedis), &categories)
		if err != nil {
			return nil, err
		}

		return categories, nil
	}

	categories := []*models.Category{}
	err = service.DB.Preload("Products").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	categoriesJSON, err := json.Marshal(categories)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, categoriesJSON, 24*time.Hour)

	return categories, nil
}

func (service *CategoryServiceImpl) DeleteCategoryById(id string) error {
	category, err := service.GetCategoryById(id)
	if err != nil {
		return err
	}

	err = service.DB.Delete(&category).Error
	if err != nil {
		return err
	}

	cacheKey := "category:" + id
	service.Redis.Del(context.Background(), cacheKey)

	return nil
}

func NewCategoryService(db *gorm.DB, redis *redis.Client) CategoryService {
	return &CategoryServiceImpl{DB: db, Redis: redis}
}
