package entity

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	ClassID uuid.UUID

	FromTime time.Time
	ToTime   time.Time
	Action   string
	Date     int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
