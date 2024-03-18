package auth

import "context"

type IUseCase interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct{}
