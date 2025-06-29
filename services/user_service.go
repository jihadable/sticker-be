package services

import (
	"github.com/jihadable/sticker-be/config"
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser() (*model.User, error)
	GetUserById() (*model.User, error)
	UpdateUserById() (*model.User, error)
	VerifyUser() (*model.User, error)
}

type UserServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *UserServiceImpl) AddUser() (*model.User, error) {
	panic("")
}

func (service *UserServiceImpl) GetUserById() (*model.User, error) {
	panic("")
}

func (service *UserServiceImpl) UpdateUserById() (*model.User, error) {
	panic("")
}

func (service *UserServiceImpl) VerifyUser() (*model.User, error) {
	panic("")
}

func NewUserService() UserService {
	return &UserServiceImpl{DB: config.DB(), Redis: config.Redis()}
}
