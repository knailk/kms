package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRequested struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Username    string
	FullName    string
	ParentName  string
	Password    string
	Email       string
	PhoneNumber string
	BirthDate   time.Time
	Gender      string
	Status      UserRequestedStatus
	ClassID     uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRequestedStatus string

var (
	UserRequestedStatusPending  UserRequestedStatus = "pending"
	UserRequestedStatusApproved UserRequestedStatus = "approved"
	UserRequestedStatusRejected UserRequestedStatus = "rejected"
)
