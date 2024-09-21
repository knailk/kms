package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserClass struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Username  string
	ClassID   uuid.UUID
	Status    string // joined, studying, complete, ...etc
	CreatedAt time.Time
	UpdatedAt time.Time

	CheckInOut []CheckInOut `gorm:"foreignKey:UserClassID"`
}

type UserClassStatus string

const (
	UserClassStatusJoined   UserClassStatus = "joined"
	UserClassStatusStudying UserClassStatus = "studying"
	UserClassStatusComplete UserClassStatus = "complete"
	UserClassStatusCanceled UserClassStatus = "canceled"
)
