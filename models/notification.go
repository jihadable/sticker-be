package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	Id      string `gorm:"column:id;primaryKey" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Message string `gorm:"column:message" json:"message"`
	Type    string `gorm:"column:type" json:"type"`

	Recipients []User `gorm:"many2many:notification_recipients;joinForeignKey:NotificationId;joinReferences:RecipientId"`
}

func (model *Notification) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
