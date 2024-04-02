package entity

import (
	"time"

	"github.com/google/uuid"
)

type CheckInOut struct {
	ID       uuid.UUID        `gorm:"primaryKey"`
	ClassID  uuid.UUID        `gorm:"index"`
	Username string           `gorm:"index"`
	Action   CheckInOutAction `gorm:"type:\"CheckInOutAction\""`
	Date     int64            `gorm:"index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CheckInOutAction string

const (
	CheckIn  CheckInOutAction = "check_in"
	CheckOut CheckInOutAction = "check_out"
)

func (c CheckInOutAction) String() string {
	return string(c)
}
