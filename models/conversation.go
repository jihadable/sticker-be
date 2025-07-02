package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	Id         uuid.UUID `gorm:"column:id;primaryKey"`
	CustomerId string    `gorm:"column:customer_id"`

	Customer *User     `gorm:"foreignKey:CustomerId;references:Id"`
	Messages []Message `gorm:"foreignKey:ConversationId;references:Id"`
}

func (model *Conversation) BeforeCreate(tx *gorm.DB) error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
	}
	return nil
}
