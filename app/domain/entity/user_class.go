package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserClass struct {
	Username  string
	ClassID   uuid.UUID
	Status    string // joined, studying, complete, ...etc
	CreatedAt time.Time
	UpdatedAt time.Time
}
