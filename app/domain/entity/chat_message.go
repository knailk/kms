package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type ChatMessage struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	ChatSessionID uuid.UUID
	Sender        string
	Message       string
	Type          MessageType `gorm:"type:\"MessageType\""`
	IsRead        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IsDeleted     soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

type MessageType string

const (
	MessageText     MessageType = "text"
	MessageImage    MessageType = "image"
	MessageVideo    MessageType = "video"
	MessageFile     MessageType = "file"
	MessageLink     MessageType = "link"
	MessageVoice    MessageType = "voice"
	MessageSticker  MessageType = "sticker"
	MessageLocation MessageType = "location"
)

func (m MessageType) String() string {
	return string(m)
}
