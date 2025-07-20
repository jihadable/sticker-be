package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRecipient struct {
	Id             string `gorm:"column:id;primaryKey" json:"id"`
	NotificationId string `gorm:"column:notification_id" json:"notification_id"`
	RecipientId    string `gorm:"column:recipient_id" json:"recipient_id"`
	IsRead         bool   `gorm:"column:is_read" json:"is_read"`

	Notification *Notification `gorm:"foreignKey:NotificationId;references:Id"`
	Recipient    *User         `gorm:"foreignKey:RecipientId;references:Id"`
}

func (model *NotificationRecipient) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
