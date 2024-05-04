package auth

import (
	"kms/app/domain/entity"
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RegisterRequest struct {
	Username    string    `json:"username"`
	FullName    string    `json:"fullName"`
	ParentName  string    `json:"parentName"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	BirthDate   time.Time `json:"birthDate"`
	Gender      string    `json:"gender"`
	ClassID     uuid.UUID `json:"classID"`
}

type RegisterListRequest struct {
	Status string `form:"status"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type RegisterConfirmRequest struct {
	ID     uuid.UUID                  `json:"id"`
	Action entity.UserRequestedStatus `json:"action"`
}
