package class

import "context"

type IUseCase interface {
	CreateClass(ctx context.Context, req *CreateClassRequest) (*CreateClassResponse, error)
	GetClass(ctx context.Context, req *GetClassRequest) (*GetClassResponse, error)
	UpdateClass(ctx context.Context, req *UpdateClassRequest) (*UpdateClassResponse, error)
	ListClasses(ctx context.Context, req *ListClassesRequest) (*ListClassesResponse, error)
	DeleteClass(ctx context.Context, req *DeleteClassRequest) (*DeleteClassResponse, error)
}
