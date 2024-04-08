package class

import "context"

type IUseCase interface {
	GetClass(ctx context.Context, req *GetClassRequest) (*GetClassResponse, error)
	ListClasses(ctx context.Context, req *ListClassesRequest) (*ListClassesResponse, error)
	CreateClass(ctx context.Context, req *CreateClassRequest) (*CreateClassResponse, error)
	UpdateClass(ctx context.Context, req *UpdateClassRequest) (*UpdateClassResponse, error)
	DeleteClass(ctx context.Context, req *DeleteClassRequest) (*DeleteClassResponse, error)
}
