package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MessageService interface {
	AddMessage(message *models.Message) (*models.Message, error)
}

type MessageServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *MessageServiceImpl) AddMessage(message *models.Message) (*models.Message, error) {
	err := service.DB.Create(message).Error
	if err != nil {
		return nil, err
	}

	return service.GetMessageById(message.Id)
}

func (service *MessageServiceImpl) GetMessageById(id string) (*models.Message, error) {
	message := models.Message{}
	err := service.DB.Where("id = ?", id).Preload("Conversation").Preload("Product").Preload("CustomProduct").Preload("Sender").First(&message).Error
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func NewMessageService(db *gorm.DB, redis *redis.Client) MessageService {
	return &MessageServiceImpl{DB: db, Redis: redis}
}
