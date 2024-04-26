package class

import "context"

type IUseCase interface {
	GetClass(ctx context.Context, req *GetClassRequest) (*GetClassResponse, error)
	ListClasses(ctx context.Context, req *ListClassesRequest) (*ListClassesResponse, error)
	CreateClass(ctx context.Context, req *CreateClassRequest) (*CreateClassResponse, error)
	UpdateClass(ctx context.Context, req *UpdateClassRequest) (*UpdateClassResponse, error)
	DeleteClass(ctx context.Context, req *DeleteClassRequest) (*DeleteClassResponse, error)

	AddMembersToClass(ctx context.Context, req *AddMembersToClassRequest) (*AddMemberToClassResponse, error)
	ListMembersInClass(ctx context.Context, req *ListMembersInClassRequest) (*ListMembersInClassResponse, error)
	RemoveMembersFromClass(ctx context.Context, req *RemoveMembersFromClassRequest) (*RemoveMembersFromClassResponse, error)

	CheckInOut(ctx context.Context, req *CheckInOutRequest) (*CheckInOutResponse, error)
	CheckInOutHistories(ctx context.Context, req *CheckInOutHistoriesRequest) (*CheckInOutHistoriesResponse, error)

	CreateSchedule(ctx context.Context, req *CreateScheduleRequest) (*CreateScheduleResponse, error)
	UpdateSchedule(ctx context.Context, req *UpdateScheduleRequest) (*UpdateScheduleResponse, error)
	DeleteSchedule(ctx context.Context, req *DeleteScheduleRequest) (*DeleteScheduleResponse, error)
}
