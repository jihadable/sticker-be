package mapper

import (
	"github.com/jihadable/sticker-be/graph/model"
	"github.com/jihadable/sticker-be/models"
)

func DBConversationTOGraphQLConversation(conversation *models.Conversation) *model.Conversation {
	messages := make([]*model.Message, len(conversation.Messages))
	for i, message := range conversation.Messages {
		messages[i] = DBMessageToGraphQLMessage(&message)
	}

	return &model.Conversation{
		ID: conversation.Id,
		Customer: func() *model.User {
			if conversation.Customer != nil {
				return DBUserToGraphQLUser(conversation.Customer)
			} else {
				return nil
			}
		}(),
		Admin: func() *model.User {
			if conversation.Admin != nil {
				return DBUserToGraphQLUser(conversation.Admin)
			} else {
				return nil
			}
		}(),
		Messages: messages,
	}
}
