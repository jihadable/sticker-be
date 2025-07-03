package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	Id              string  `gorm:"column:id;primaryKey"`
	ConversationId  string  `gorm:"column:conversation_id"`
	ProductId       *string `gorm:"column:product_id"`
	CustomProductId *string `gorm:"column:custom_product_id"`
	SenderId        string  `gorm:"column:sender_id"`
	Message         string  `gorm:"column:message"`

	Conversation  *Conversation  `gorm:"foreignKey:ConversationId;references:Id"`
	Product       *Product       `gorm:"foreignKey:ProductId;references:Id"`
	CustomProduct *CustomProduct `gorm:"foreignKey:CustomProductId;references:Id"`
	Sender        *User          `gorm:"foreignKey:SenderId;references:Id"`
}

func (model *Message) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
