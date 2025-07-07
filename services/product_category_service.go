package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductCategoryService interface {
	AddProductCategory(productCategory *models.ProductCategory) (*models.ProductCategory, error)
	DeleteProductCategory(product_id, category_id string) error
}

type ProductCategoryServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *ProductCategoryServiceImpl) AddProductCategory(productCategory *models.ProductCategory) (*models.ProductCategory, error) {
	err := service.DB.Create(&productCategory).Error
	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

func (service *ProductCategoryServiceImpl) GetProductCategory(product_id, category_id string) (*models.ProductCategory, error) {
	productCategory := models.ProductCategory{}
	err := service.DB.Where("product_id = ? AND category_id = ?", product_id, category_id).First(&productCategory).Error
	if err != nil {
		return nil, err
	}

	return &productCategory, nil
}

func (service *ProductCategoryServiceImpl) DeleteProductCategory(product_id, category_id string) error {
	productCategory, err := service.GetProductCategory(product_id, category_id)
	if err != nil {
		return err
	}

	err = service.DB.Delete(productCategory).Error

	return err
}

func NewProductCategoryService(db *gorm.DB, redis *redis.Client) ProductCategoryService {
	return &ProductCategoryServiceImpl{DB: db, Redis: redis}
}
