package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type NotificationRecipientService interface {
	AddNotificationRecipients(notificationId string, recipientIds ...string) (*models.NotificationRecipient, error)
}

type NotificationRecipientServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *NotificationRecipientServiceImpl) AddNotificationRecipients(notificationId string, recipientIds ...string) (*models.NotificationRecipient, error) {
	panic("")
}

func NewNotificationRecipientService(db *gorm.DB, redis *redis.Client) NotificationRecipientService {
	return &NotificationRecipientServiceImpl{DB: db, Redis: redis}
}
