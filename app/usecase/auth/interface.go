package auth

import (
	"context"
)

type IUseCase interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error)

	Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	// RegisterList(ctx context.Context, req *RegisterListRequest) (*RegisterListResponse, error)
	RegisterConfirm(ctx context.Context, req *RegisterConfirmRequest) (*RegisterConfirmResponse, error)
}
