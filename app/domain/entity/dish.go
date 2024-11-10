package entity

import (
	"time"

	"github.com/google/uuid"
)

type Dish struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	DayOfWeek string    `gorm:"size:255"`
	Date      int       `gorm:"index:date_unique,unique"`

	// Fields for each meal
	Breakfast      string `gorm:"size:255"`
	EatLightly     string `gorm:"size:255"`
	Lunch          string `gorm:"size:255"`
	AfternoonSnack string `gorm:"size:255"`
	Dinner         string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
