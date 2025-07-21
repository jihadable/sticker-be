package mapper

import "github.com/jihadable/sticker-be/models"

func NotificationMapper(notification *models.Notification) map[string]any {
	return map[string]any{
		"type":         notification.Type,
		"recipient_id": notification.RecipientId,
		"title":        notification.Title,
		"message":      notification.Message,
		"is_read":      notification.IsRead,
	}
}
