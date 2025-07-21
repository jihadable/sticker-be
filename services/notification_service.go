package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type NotificationService interface {
	AddNotification(notification *models.Notification, recipientIds ...string) (*models.Notification, error)
	GetNotificationById(id string) (*models.Notification, error)
}

type NotificationServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
	NotificationRecipientService
}

func (service *NotificationServiceImpl) AddNotification(notification *models.Notification, recipientIds ...string) (*models.Notification, error) {
	err := service.DB.Create(notification).Error
	if err != nil {
		return nil, err
	}

	_, err = service.NotificationRecipientService.AddNotificationRecipients(notification.Id, recipientIds...)
	if err != nil {
		return nil, err
	}

	return service.GetNotificationById(notification.Id)
}

func (service *NotificationServiceImpl) GetNotificationById(id string) (*models.Notification, error) {
	ctx := context.Background()
	cacheKey := "notification:" + id

	notificationInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && notificationInRedis != "" {
		notification := models.Notification{}
		err := json.Unmarshal([]byte(notificationInRedis), &notification)
		if err != nil {
			return nil, err
		}

		return &notification, nil
	}

	notification := models.Notification{}
	err = service.DB.Where("id = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, notificationJSON, 5*time.Minute)

	return &notification, nil
}

func NewNotificationService(db *gorm.DB, redis *redis.Client, notificationRecipientService NotificationRecipientService) NotificationService {
	return &NotificationServiceImpl{DB: db, Redis: redis, NotificationRecipientService: notificationRecipientService}
}
