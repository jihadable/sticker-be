package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CartProductService interface {
	AddCartProduct(cartProduct *models.CartProduct) (*models.CartProduct, error)
	UpdateCartProductById(id string, updatedCartProduct *models.CartProduct) (*models.CartProduct, error)
	DeleteCartProductById(id string) error
}

type CartProductServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *CartProductServiceImpl) AddCartProduct(cartProduct *models.CartProduct) (*models.CartProduct, error) {
	err := service.DB.Create(cartProduct).Error
	if err != nil {
		return nil, err
	}

	return service.GetCartProductById(cartProduct.Id)
}

func (service *CartProductServiceImpl) GetCartProductById(id string) (*models.CartProduct, error) {
	cartProduct := models.CartProduct{}
	err := service.DB.Where("id = ?", id).Preload("Cart").Preload("Product").Preload("CustomProduct").First(&cartProduct).Error
	if err != nil {
		return nil, err
	}

	return &cartProduct, nil
}

func (service *CartProductServiceImpl) UpdateCartProductById(id string, updatedCartProduct *models.CartProduct) (*models.CartProduct, error) {
	cartProduct, err := service.GetCartProductById(id)
	if err != nil {
		return nil, err
	}

	cartProduct.Quantity = updatedCartProduct.Quantity
	cartProduct.Size = updatedCartProduct.Size

	err = service.DB.Save(cartProduct).Error
	if err != nil {
		return nil, err
	}

	return service.GetCartProductById(cartProduct.Id)
}

func (service *CartProductServiceImpl) DeleteCartProductById(id string) error {
	cartProduct, err := service.GetCartProductById(id)
	if err != nil {
		return err
	}

	err = service.DB.Delete(cartProduct).Error

	return err
}

func NewCartProductService(db *gorm.DB, redis *redis.Client) CartProductService {
	return &CartProductServiceImpl{DB: db, Redis: redis}
}
