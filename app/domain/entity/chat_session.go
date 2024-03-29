package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type ChatSession struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string

	ChatParticipants []ChatParticipant `gorm:"foreignKey:ChatSessionID"`
	ChatMessages     []ChatMessage     `gorm:"foreignKey:ChatSessionID"`

	CreatedAt time.Time
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
