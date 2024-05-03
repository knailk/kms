package auth

import (
	"kms/app/domain/entity"
	"kms/pkg/authjwt"
	"time"

	"github.com/google/uuid"
)

type LoginResponse struct {
	Username     string        `json:"username"`
	Email        string        `json:"email"`
	Role         string        `json:"role"`
	ParentName   string        `json:"parentFullName"`
	FullName     string        `json:"fullName"`
	Gender       string        `json:"gender"`
	PhoneNumber  string        `json:"phoneNumber"`
	BirthDate    time.Time     `json:"birthDate"`
	PictureURL   string        `json:"pictureURL"`
	Address      string        `json:"address"`
	Longitude    *float64      `json:"longitude"`
	Latitude     *float64      `json:"latitude"`
	CreatedAt    time.Time     `json:"createdAt"`
	AccessToken  authjwt.Token `json:"-"`
	RefreshToken authjwt.Token `json:"-"`
}

type RefreshResponse struct {
	AccessToken  authjwt.Token `json:"-"`
	RefreshToken authjwt.Token `json:"-"`
}

type RegisterResponse struct {
}

type RegisterListResponse struct {
}

type RegisterConfirmResponse struct {
	ClassID  uuid.UUID
	Username string
	Status   entity.UserRequestedStatus
}

type UserInfo struct{}
