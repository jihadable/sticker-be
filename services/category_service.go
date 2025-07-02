package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CategoryService interface {
	AddCategory()
	GetCategoryById()
	GetCategoriesByProduct()
	GetCategories()
	DeleteCategoryById()
}

type CategoryServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CategoryServiceImpl) AddCategory() {

}

func (service *CategoryServiceImpl) GetCategoryById() {

}

func (service *CategoryServiceImpl) GetCategoriesByProduct() {

}

func (service *CategoryServiceImpl) GetCategories() {

}

func (service *CategoryServiceImpl) DeleteCategoryById() {

}

func NewCategoryService(db *gorm.DB, redis *redis.Client) CategoryService {
	return &CategoryServiceImpl{DB: db, Redis: redis}
}
