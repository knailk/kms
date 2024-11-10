package dish

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/logger"

	"github.com/google/uuid"
)

type useCase struct {
	repo *repository.PostgresRepository
}

func NewUseCase(repo *repository.PostgresRepository) IUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) CreateDish(ctx context.Context, req *CreateDishRequest) (*CreateDishResponse, error) {
	const op errs.Op = "useCase.dish.CreateDish"

	errKind := req.Validate()
	if errKind != errs.Other {
		return nil, errs.E(op, errKind, "validate request error")
	}

	dishID := uuid.New()

	err := uc.repo.Dish.Create(&entity.Dish{
		ID:             dishID,
		DayOfWeek:      req.DayOfWeek,
		Date:           req.Date,
		Breakfast:      req.Breakfast,
		EatLightly:     req.EatLightly,
		Lunch:          req.Lunch,
		AfternoonSnack: req.AfternoonSnack,
		Dinner:         req.Dinner,
	})
	if err != nil {
		logger.Error(op, " create dish error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &CreateDishResponse{}, nil
}

func (uc *useCase) UpdateDish(ctx context.Context, req *UpdateDishRequest) (*UpdateDishResponse, error) {
	const op errs.Op = "useCase.dish.UpdateDish"

	errKind := req.Validate()
	if errKind != errs.Other {
		return nil, errs.E(op, errKind, "validate request error")
	}

	_, err := uc.repo.Dish.Where(uc.repo.Dish.ID.Eq(req.DishID)).Updates(
		&entity.Dish{
			Breakfast:      req.Breakfast,
			EatLightly:     req.EatLightly,
			Lunch:          req.Lunch,
			AfternoonSnack: req.AfternoonSnack,
			Dinner:         req.Dinner,
		},
	)
	if err != nil {
		logger.Error(op, " update dish error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &UpdateDishResponse{}, nil
}

func (uc *useCase) DeleteDish(ctx context.Context, req *DeleteDishRequest) error {
	const op errs.Op = "useCase.dish.DeleteDish"

	_, err := uc.repo.Dish.Where(uc.repo.Dish.ID.Eq(req.DishID)).Delete()
	if err != nil {
		logger.Error(op, " delete dish error :", err)
		return errs.E(op, errs.Database, err)
	}

	return nil
}

func (uc *useCase) GetDishesForWeek(ctx context.Context, req *GetDishesForWeekRequest) (*GetDishesForWeekResponse, error) {
	const op errs.Op = "useCase.dish.GetDishesForWeek"

	dishes, err := uc.repo.Dish.
		Where(uc.repo.Dish.Date.Lte(req.ToDate), uc.repo.Dish.Date.Gte(req.FromDate)).
		Order(uc.repo.Dish.Date).
		Find()
	if err != nil {
		logger.Error(op, " get dishes for week error :", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &GetDishesForWeekResponse{
		Dishes: toDishesResponse(dishes),
	}, nil
}
