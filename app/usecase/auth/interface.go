package auth

import (
	"context"
)

type IUseCase interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	GetInfo(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error)
	Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error)
}
