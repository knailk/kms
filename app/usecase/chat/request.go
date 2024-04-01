package chat

import (
	"kms/app/domain/entity"
	"kms/app/errs"

	"github.com/google/uuid"
)

type CreateChatRequest struct {
	Owner        string   `json:"-"`
	Participants []string `json:"participants"`
}

func (c *CreateChatRequest) Validate() errs.Kind {
	if len(c.Participants) < 1 {
		return errs.InvalidRequest
	}

	return errs.Other
}

type AddMemberRequest struct {
	Adder         string    `json:"-"`
	UserID        string    `json:"user_id"`
	ChatSessionID uuid.UUID `json:"chat_id"`
}

type ListChatsRequest struct {
	UserRequester string `json:"-"`
}

type GetChatRequest struct {
	UserRequester string    `json:"-"`
	ChatSessionID uuid.UUID `json:"-"`
}

type UpdateChatRequest struct {
	ChatSessionID uuid.UUID `json:"-"`
	Name          string    `json:"name"`
}

type DeleteChatRequest struct {
	ChatSessionID uuid.UUID `json:"-"`
}

type CreateMessageRequest struct {
	Sender string `json:"-"`

	ChatSessionID uuid.UUID          `json:"-"`
	Message       string             `json:"message"`
	Type          entity.MessageType `json:"type"`
}

func (c *CreateMessageRequest) Validate() errs.Kind {
	if c.Message == "" || c.Type == "" || c.ChatSessionID == uuid.Nil || c.Sender == "" {
		return errs.InvalidRequest
	}

	return errs.Other
}
