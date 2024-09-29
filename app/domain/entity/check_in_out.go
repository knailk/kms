package entity

import (
	"time"

	"github.com/google/uuid"
)

type CheckInOut struct {
	ID          uuid.UUID        `gorm:"primaryKey"`
	UserClassID uuid.UUID        `gorm:"index:user_date_unique,unique"`
	Action      CheckInOutAction `gorm:"type:\"CheckInOutAction\""`
	Date        int64            `gorm:"index:user_date_unique,unique"`

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
