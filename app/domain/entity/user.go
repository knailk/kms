package entity

import (
	"time"
)

type User struct {
	// Username: The unique identifier for the User's profile
	Username string `json:"username" gorm:"primaryKey"`

	// Password: The User's password
	Password string `json:"password"`

	Role UserRole `json:"role" gorm:"type:\"UserRole\""`

	// FullName: The person's full name.
	FullName string `json:"fullName"`

	// Gender: The user's gender.
	Gender string `json:"gender"`

	// Email: The primary email for the User
	Email string `json:"email" gorm:"unique"`

	// BirthDate: The full birthDate of a person (e.g. Dec 18, 1953)
	BirthDate time.Time `json:"birthDate"`

	// PhoneNumber: The phone number of the person.
	PhoneNumber string `json:"phoneNumber"`

	// PictureURL: URL of the person's picture image for the profile.
	PictureURL string `json:"pictureURL"`

	// Address: The person's address.
	Address string `json:"address"`

	// CreatedAt: The time the User was created.
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;default:now()"`

	// UpdatedAt: The time the User was last updated.
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;default:current_timestamp"`

	// IsDeleted: Whether the User is deleted or not.
	IsDeleted bool `json:"isDeleted" gorm:"default:false"`
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
