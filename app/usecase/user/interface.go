package user

import "context"

type IUseCase interface {
	GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error)
	SearchUser(ctx context.Context, req *SearchUserRequest) (*SearchUserResponse, error)
}
