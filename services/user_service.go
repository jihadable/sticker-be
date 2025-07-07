package services

import (
	"errors"

	"github.com/jihadable/sticker-be/models"
	"github.com/jihadable/sticker-be/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(user *models.User) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUserById(id string, updatedUser *models.User) (*models.User, error)
	VerifyUser(email string, password string) (*models.User, error)
}

type UserServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *UserServiceImpl) AddUser(user *models.User) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	err = service.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	user := models.User{}

	err := service.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserServiceImpl) UpdateUserById(id string, updatedUser *models.User) (*models.User, error) {
	user := models.User{}

	err := service.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	user.Name = updatedUser.Name
	user.Phone = updatedUser.Phone
	user.Address = updatedUser.Address

	err = service.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserServiceImpl) VerifyUser(email, password string) (*models.User, error) {
	user := models.User{}

	err := service.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("")
	}

	return &user, nil
}

func NewUserService(db *gorm.DB, redis *redis.Client) UserService {
	return &UserServiceImpl{DB: db, Redis: redis}
}
