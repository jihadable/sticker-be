package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       string `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
	Phone    string `gorm:"column:phone"`
	Address  string `gorm:"column:address"`

	CustomProducts []CustomProduct `gorm:"foreignKey:CustomerId;references:Id"`
	Cart           Cart            `gorm:"foreignKey:CustomerId;references:Id"`
	Orders         []Order         `gorm:"foreignKey:CustomerId;references:Id"`
	Conversation   Conversation    `gorm:"foreignKey:CustomerId;references:Id"`
}

func (model *User) BeforeCreate(tx *gorm.DB) error {
	if model.Id == "" {
		model.Id = uuid.NewString()
	}
	if model.Role == "" {
		model.Role = "customer"
	}

	return nil
}
