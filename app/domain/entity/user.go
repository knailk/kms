package entity

import (
	"time"
)

type User struct {
	// Username: The unique identifier for the User's profile
	Username string `gorm:"primaryKey"`

	// Password: The User's password
	Password string

	// Role: The User's role in system
	Role UserRole `gorm:"type:\"UserRole\""`

	// FullName: The person's full name.
	FullName string

	// Gender: The user's gender.
	Gender string

	// Email: The primary email for the User
	Email string `gorm:"unique"`

	// BirthDate: The full birthDate of a person (e.g. Dec 18, 1953)
	BirthDate time.Time

	// PhoneNumber: The phone number of the person.
	PhoneNumber string

	// PictureURL: URL of the person's picture image for the profile.
	PictureURL string

	// Address: The person's address.
	Address string

	// CreatedAt: The time the User was created.
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`

	// UpdatedAt: The time the User was last updated.
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`

	// IsDeleted: Whether the User is deleted or not.
	IsDeleted bool `gorm:"default:false"`
}

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleStudent UserRole = "student"
	UserRoleTeacher UserRole = "teacher"
	UserRoleChef    UserRole = "chef"
	UserRoleDriver  UserRole = "driver"
)

type UserRoleList []UserRole
