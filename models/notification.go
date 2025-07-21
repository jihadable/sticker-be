package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	Id          string `gorm:"column:id;primaryKey" json:"id"`
	Type        string `gorm:"column:type" json:"type"`
	RecipientId string `grom:"column:recipient_id" json:"recipient_id"`
	Title       string `gorm:"column:title" json:"title"`
	Message     string `gorm:"column:message" json:"message"`
	IsRead      bool   `gorm:"column:is_read;default:false" json:"is_read"`

	Recipient *User `gorm:"foreignKey:RecipientId;references:Id"`
}

func (model *Notification) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
