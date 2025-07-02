package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser()
	GetUserById()
	UpdateUserById()
	VerifyUser()
}

type UserServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *UserServiceImpl) AddUser() {

}

func (service *UserServiceImpl) GetUserById() {

}

func (service *UserServiceImpl) UpdateUserById() {

}

func (service *UserServiceImpl) VerifyUser() {

}

func NewUserService(db *gorm.DB, redis *redis.Client) UserService {
	return &UserServiceImpl{DB: db, Redis: redis}
}
