package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ConversationService interface {
	AddConversation(conversation *models.Conversation) (*models.Conversation, error)
	GetConversationByCustomer(customer_id string) (*models.Conversation, error)
}

type ConversationServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *ConversationServiceImpl) AddConversation(conversation *models.Conversation) (*models.Conversation, error) {
	err := service.DB.Create(conversation).Error
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (service *ConversationServiceImpl) GetConversationById(id string) (*models.Conversation, error) {
	ctx := context.Background()
	cacheKey := "conversation:" + id

	conversationInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && conversationInRedis != "" {
		conversation := models.Conversation{}
		err = json.Unmarshal([]byte(conversationInRedis), &conversation)
		if err != nil {
			return nil, err
		}

		return &conversation, nil
	}

	conversation := models.Conversation{}
	err = service.DB.Where("id = ?", id).Preload("Customer").Preload("Messages").First(&conversation).Error
	if err != nil {
		return nil, err
	}

	conversationJSON, err := json.Marshal(conversation)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, conversationJSON, 30*time.Minute)

	return &conversation, nil
}

func (service *ConversationServiceImpl) GetConversationByCustomer(customer_id string) (*models.Conversation, error) {
	conversation := models.Conversation{}
	err := service.DB.Where("customer_id = ?", customer_id).Preload("Customer").Preload("Messages").First(&conversation).Error
	if err != nil {
		return nil, err
	}

	return &conversation, nil
}

func NewConversationService(db *gorm.DB, redis *redis.Client) ConversationService {
	return &ConversationServiceImpl{DB: db, Redis: redis}
}
