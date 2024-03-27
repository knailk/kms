package auth

import (
	"context"
)

type IUseCase interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error)
	GetProfile(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error)
	UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (*UpdateProfileResponse, error)
}
