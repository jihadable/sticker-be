package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBUserToGraphQLUser(user *models.User) *model.User {
	return &model.User{}
}
