package dish

import "context"

type IUseCase interface {
	CreateDish(ctx context.Context, req *CreateDishRequest) (*CreateDishResponse, error)
	UpdateDish(ctx context.Context, req *UpdateDishRequest) (*UpdateDishResponse, error)
	DeleteDish(ctx context.Context, req *DeleteDishRequest) error
	GetDishesForWeek(ctx context.Context, req *GetDishesForWeekRequest) (*GetDishesForWeekResponse, error)
}
