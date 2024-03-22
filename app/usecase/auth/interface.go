package auth

import (
	"context"
	"kms/pkg/authjwt"
	"time"
)

type IUseCase interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username     string        `json:"username"`
	Email        string        `json:"email"`
	Role         string        `json:"role"`
	FullName     string        `json:"fullName"`
	Gender       string        `json:"gender"`
	PhoneNumber  string        `json:"phoneNumber"`
	BirthDate    time.Time     `json:"birthDate"`
	PictureURL   string        `json:"pictureURL"`
	Address      string        `json:"address"`
	AccessToken  authjwt.Token `json:"-"`
	RefreshToken authjwt.Token `json:"-"`
}
