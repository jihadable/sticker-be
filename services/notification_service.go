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
	AddNotification(notification *models.Notification) (*models.Notification, error)
	GetNotificationById(id string) (*models.Notification, error)
	GetNotificationsByRecipient(recipient_id string) ([]*models.Notification, error)
	ReadNotificationById(id string) error
	ReadAllNotificationsByRecipient(recipient_id string) error
}

type NotificationServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *NotificationServiceImpl) AddNotification(notification *models.Notification) (*models.Notification, error) {
	err := service.DB.Create(notification).Error
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

func (service *NotificationServiceImpl) GetNotificationsByRecipient(recipient_id string) ([]*models.Notification, error) {
	ctx := context.Background()
	cacheKey := "notification:recipient:" + recipient_id

	notificationsInRedis, err := service.Redis.Get(ctx, cacheKey).Result()
	if err == nil && notificationsInRedis != "" {
		notifications := []*models.Notification{}
		err = json.Unmarshal([]byte(notificationsInRedis), &notifications)
		if err != nil {
			return nil, err
		}

		return notifications, nil
	}

	notifications := []*models.Notification{}
	err = service.DB.Where("recipient_id = ?", recipient_id).Preload("Recipient").Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	notificationsJSON, err := json.Marshal(notifications)
	if err != nil {
		return nil, err
	}
	service.Redis.Set(ctx, cacheKey, notificationsJSON, 5*time.Minute)

	return notifications, nil
}

func (service *NotificationServiceImpl) ReadNotificationById(id string) error {
	notification, err := service.GetNotificationById(id)
	if err != nil {
		return err
	}

	notification.IsRead = true

	err = service.DB.Save(notification).Error
	if err != nil {
		return err
	}

	cacheKeys := []string{"notification:" + notification.Id, "notification:recipient:" + notification.RecipientId}
	service.Redis.Del(context.Background(), cacheKeys...)

	return nil
}

func (service *NotificationServiceImpl) ReadAllNotificationsByRecipient(recipient_id string) error {
	notifications, err := service.GetNotificationsByRecipient(recipient_id)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		notification.IsRead = true
	}

	err = service.DB.Save(notifications).Error
	if err != nil {
		return err
	}

	cacheKeys := []string{}
	for _, notification := range notifications {
		cacheKeys = append(cacheKeys, "notification:", notification.Id)
	}
	cacheKeys = append(cacheKeys, "notification:recipient:", recipient_id)
	service.Redis.Del(context.Background(), cacheKeys...)

	return nil
}

func NewNotificationService(db *gorm.DB, redis *redis.Client) NotificationService {
	return &NotificationServiceImpl{DB: db, Redis: redis}
}
