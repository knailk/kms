package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type ChatParticipant struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	ChatSessionID uuid.UUID
	Username      string
	IsOwner       bool
	CreatedAt     time.Time
	IsDeleted     soft_delete.DeletedAt `gorm:"softDelete:flag"`

	User User `gorm:"foreignKey:username"`
}
