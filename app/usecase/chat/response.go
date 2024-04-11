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
	ID            uuid.UUID        `json:"id"`
	Name          string           `json:"name"`
	ChatPicture   string           `json:"chatPicture"`
	ChatMessages  []MessageByDate  `json:"chatMessages"`
	LatestMessage *MessageResponse `json:"latestMessage"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MessageByDate struct {
	Date     string             `json:"date"`
	Messages []*MessageResponse `json:"messages"`
}

type MessageResponse struct {
	ID             uuid.UUID       `json:"id"`
	Message        string          `json:"content"`
	Type           string          `json:"type"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"-"`
	SenderResponse *SenderResponse `json:"sender"`
}

type SenderResponse struct {
	Username   string `json:"username"`
	SenderName string `json:"name"`
	Avatar     string `json:"avatar"`
}

type DeleteChatResponse struct{}

type CreateMessageResponse struct{}
