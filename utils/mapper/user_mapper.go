package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBUserToGraphQLUser(user *models.User) *model.User {
	return &model.User{
		ID:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Role:    model.Role(user.Role),
		Phone:   user.Phone,
		Address: user.Address,
	}
}
