package chat

import (
	"kms/app/errs"

	"github.com/google/uuid"
)

type CreateChatRequest struct {
	Owner        string   `json:"-"`
	Participants []string `json:"participants"`
}

func (c *CreateChatRequest) Validate() errs.Kind {
	if len(c.Participants) < 2 {
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
	ChatID uuid.UUID `json:"-"`
	Name   string    `json:"name"`
}

type DeleteChatRequest struct {
	ChatID uuid.UUID `json:"-"`
}
