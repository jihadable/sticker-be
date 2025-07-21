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
		ID:       conversation.Id,
		Customer: DBUserToGraphQLUser(conversation.Customer),
		Admin:    DBUserToGraphQLUser(conversation.Admin),
		Messages: messages,
	}
}
