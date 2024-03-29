package chat

import (
	"kms/app/domain/entity"

	"github.com/google/uuid"
)

type CreateChatRequest struct {
	Owner        string             `json:"-"`
	Name         string             `json:"name"`
	Participants []string           `json:"participants"`
	Message      string             `json:"message"`
	MessageType  entity.MessageType `json:"message_type"`
}

type AddMemberRequest struct {
	UserID        string    `json:"user_id"`
	ChatSessionID uuid.UUID `json:"chat_id"`
}

type ListChatsRequest struct {
	Owner string `json:"-"`
}

type UpdateChatRequest struct {
	ChatID uuid.UUID `json:"chat_id"`
	Name   string    `json:"name"`
}

type GetChatRequest struct {
	ChatID uuid.UUID `json:"chat_id"`
}
