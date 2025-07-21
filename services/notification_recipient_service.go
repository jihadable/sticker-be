package services

import (
	"github.com/jihadable/sticker-be/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type NotificationRecipientService interface {
	AddNotificationRecipients(notificationId string, recipientIds ...string) ([]models.NotificationRecipient, error)
}

type NotificationRecipientServiceImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (service *NotificationRecipientServiceImpl) AddNotificationRecipients(notificationId string, recipientIds ...string) ([]models.NotificationRecipient, error) {
	notificationRecipients := make([]models.NotificationRecipient, len(recipientIds))

	for i, recipientId := range recipientIds {
		notificationRecipients[i] = models.NotificationRecipient{
			NotificationId: notificationId,
			RecipientId:    recipientId,
			IsRead:         false,
		}
	}

	err := service.DB.Create(&notificationRecipients).Error
	if err != nil {
		return nil, err
	}

	return notificationRecipients, nil
}

func NewNotificationRecipientService(db *gorm.DB, redis *redis.Client) NotificationRecipientService {
	return &NotificationRecipientServiceImpl{DB: db, Redis: redis}
}
