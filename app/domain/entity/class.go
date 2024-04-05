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

	Schedules []Schedule `gorm:"foreignKey:ClassID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
