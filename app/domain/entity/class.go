package entity

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	TeacherID string    `gorm:"index"`
	DriverID  string
	FromDate  int64
	ToDate    int64
	ClassName string
	AgeGroup  int
	Price     float64
	Currency  string // VND, USD, EUR, ...

	Schedules []Schedule `gorm:"foreignKey:ClassID"`

	User []UserClass `gorm:"foreignKey:ClassID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
