package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserClass struct {
	Username  string    `gorm:"primaryKey"`
	ClassID   uuid.UUID `gorm:"primaryKey"`
	Status    string    // joined, studying, complete, ...etc
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserClassStatus string

const (
	UserClassStatusJoined    UserClassStatus = "joined"
	UserClassStatusStudying  UserClassStatus = "studying"
	UserClassStatusComplete  UserClassStatus = "complete"
	UserClassStatusCancelled UserClassStatus = "cancelled"
)
