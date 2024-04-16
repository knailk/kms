package user

import (
	"time"
)

type GetUserResponse struct {
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role,omitempty"`
	FullName    string    `json:"fullName"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phoneNumber"`
	BirthDate   time.Time `json:"birthDate"`
	PictureURL  string    `json:"pictureURL"`
	Address     string    `json:"address"`
	Longitude   *float64  `json:"longitude"`
	Latitude    *float64  `json:"latitude"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UpdateUserResponse struct{}

type SearchUserResponse struct {
	Users []*GetUserResponse `json:"users"`
}
