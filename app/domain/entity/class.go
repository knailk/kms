package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Class struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	TeacherID   string
	DriverID    string
	FromDate    int64
	ToDate      int64
	Description string
	Status      ClassStatus
	ClassName   string
	AgeGroup    int
	Price       float64
	Currency    string // VND, USD, EUR, ...

	Schedules []Schedule `gorm:"foreignKey:ClassID"`

	User []UserClass `gorm:"foreignKey:ClassID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

type ClassStatus int

const (
	// Enum values for CourseStatus
	Scheduled ClassStatus = iota
	InProgress
	Completed
	Cancelled
)

// String returns the string representation of the CourseStatus enum
func (status ClassStatus) String() string {
	return [...]string{
		"Scheduled",
		"In Progress",
		"Completed",
		"Cancelled",
	}[status]
}
