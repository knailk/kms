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
	Status      string
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
