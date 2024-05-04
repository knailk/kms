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
	ParentName   string        `json:"parentName"`
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
	RegisterList []*UserRequestedResponse `json:"registerList"`
}

type UserRequestedResponse struct {
	ID          uuid.UUID                  `json:"id"`
	Username    string                     `json:"username"`
	FullName    string                     `json:"fullName"`
	ParentName  string                     `json:"parentName"`
	Email       string                     `json:"email"`
	PhoneNumber string                     `json:"phoneNumber"`
	BirthDate   time.Time                  `json:"birthDate"`
	Gender      string                     `json:"gender"`
	Status      entity.UserRequestedStatus `json:"status"`
	ClassID     uuid.UUID                  `json:"classID"`
	CreatedAt   time.Time                  `json:"createdAt"`

	Class *ClassResponse `json:"class"`
}

type ClassResponse struct {
	ClassName string `json:"className"`
}

type RegisterConfirmResponse struct {
	ClassID  uuid.UUID
	Username string
	Status   entity.UserRequestedStatus
}

type UserInfo struct{}
