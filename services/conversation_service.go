package services

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ConversationService interface {
	AddConversation()
	GetConversationByUser()
}

type ConversationServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *ConversationServiceImpl) AddConversation() {

}

func (service *ConversationServiceImpl) GetConversationByUser() {

}

func NewConversationService(db *gorm.DB, redis *redis.Client) ConversationService {
	return &ConversationServiceImpl{DB: db, Redis: redis}
}
