package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MessageService interface {
	AddMessage()
}

type MessageServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *MessageServiceImpl) AddMessage() {

}

func NewMessageService(db *gorm.DB, redis *redis.Client) MessageService {
	return &MessageServiceImpl{DB: db, Redis: redis}
}
