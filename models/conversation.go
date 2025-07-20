package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	Id         string `gorm:"column:id;primaryKey"`
	CustomerId string `gorm:"column:customer_id"`
	AdminId    string `gorm:"admin_id"`

	Customer *User     `gorm:"foreignKey:CustomerId;references:Id"`
	Admin    *User     `gorm:"foreignKey:AdminId;references:Id"`
	Messages []Message `gorm:"foreignKey:ConversationId;references:Id"`
}

func (model *Conversation) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
