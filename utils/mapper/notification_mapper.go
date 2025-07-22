package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBNotificationToGraphQLNotification(notification *models.Notification) *model.Notification {
	return &model.Notification{
		ID:      notification.Id,
		Type:    notification.Type,
		Title:   notification.Title,
		Message: notification.Message,
		IsRead:  notification.IsRead,
		Recipient: func() *model.User {
			if notification.Recipient != nil {
				return DBUserToGraphQLUser(notification.Recipient)
			} else {
				return nil
			}
		}(),
	}
}
