package chat

import (
	"time"

	"github.com/google/uuid"
)

type CreateChatResponse struct{}

type AddMemberResponse struct {
	IsNewChat bool `json:"isNewChat"`
}

type ListChatsResponse struct {
	ChatSessions []*GetChatResponse `json:"chatSessions"`
}

type UpdateChatResponse struct{}

type GetChatResponse struct {
	ID           uuid.UUID          `json:"id"`
	Name         string             `json:"name"`
	ChatMessages []*MessageResponse `json:"chatMessages,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}

type MessageResponse struct {
	ID      uuid.UUID `json:"id"`
	Sender  string    `json:"sender"`
	Message string    `json:"message"`
	Type    string    `json:"type"`

	SenderName string    `json:"senderName"`
	PictureURL string    `json:"pictureURL"`
	CreatedAt  time.Time `json:"createdAt"`
}

type DeleteChatResponse struct{}
