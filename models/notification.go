package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	Id          string `gorm:"column:id;primaryKey" json:"id"`
	UserId      string `gorm:"column:user_id" json:"user_id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`

	User *User `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

func (model *Notification) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	return nil
}
