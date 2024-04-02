package check_in_out

import "context"

type IUseCase interface {
	CheckInOut(ctx context.Context, req CheckInOutRequest) (CheckInOutResponse, error)
	ListUsersInClass(ctx context.Context, req ListUsersInClassRequest) (ListUsersInClassResponse, error)
}
