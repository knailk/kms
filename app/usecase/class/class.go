package class

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/logger"

	"github.com/google/uuid"
	"gorm.io/gen"
)

type useCase struct {
	repo *repository.PostgresRepository
}

func NewUseCase(repo *repository.PostgresRepository) IUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) CreateClass(ctx context.Context, req *CreateClassRequest) (*CreateClassResponse, error) {
	const op errs.Op = "class.useCase.CreateClass"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	schedules := make([]entity.Schedule, 0)
	for _, s := range req.Schedules {
		schedules = append(schedules, entity.Schedule{
			FromTime: s.FromTime,
			ToTime:   s.ToTime,
			Action:   s.Action,
			Date:     s.Date,
		})
	}

	err := uc.repo.Class.Create(&entity.Class{
		ID:        uuid.New(),
		TeacherID: req.TeacherID,
		DriverID:  req.DriverID,
		FromDate:  req.FromDate,
		ToDate:    req.ToDate,
		ClassName: req.ClassName,
		AgeGroup:  req.AgeGroup,
		Price:     req.Price,
		Currency:  req.Currency,
		Schedules: schedules,
	})
	if err != nil {
		logger.Error(op, " create class error: ", err)
		return nil, errs.E(op, errs.Database, "create class error")
	}

	return &CreateClassResponse{}, nil
}

func (uc *useCase) GetClass(ctx context.Context, req *GetClassRequest) (*GetClassResponse, error) {
	const op errs.Op = "class.useCase.GetClass"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	filter := make([]gen.Condition, 0)

	if req.ID != uuid.Nil {
		filter = append(filter, uc.repo.Class.ID.Eq(req.ID))
	}

	if req.ClassName != "" {
		filter = append(filter, uc.repo.Class.ClassName.Eq(req.ClassName))
	}

	class, err := uc.repo.Class.Where(filter...).
		Preload(uc.repo.Class.Schedules.On(
			uc.repo.Schedule.Date.Between(req.FromDate, req.ToDate))).
		First()
	if err != nil {
		logger.Error(op, " get class error: ", err)
		return nil, errs.E(op, errs.Database, "get class error")
	}

	return toGetClassResponse(class), nil
}

func (uc *useCase) UpdateClass(ctx context.Context, req *UpdateClassRequest) (*UpdateClassResponse, error) {
	return nil, nil
}

func (uc *useCase) ListClasses(ctx context.Context, req *ListClassesRequest) (*ListClassesResponse, error) {
	const op errs.Op = "class.useCase.ListClasses"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	filter := make([]gen.Condition, 0)

	if req.FromDate > 0 {
		filter = append(filter, uc.repo.Class.FromDate.Gte(req.FromDate))
	}

	if req.ToDate > 0 {
		filter = append(filter, uc.repo.Class.ToDate.Lte(req.ToDate))
	}

	if req.AgeGroup != 0 {
		filter = append(filter, uc.repo.Class.AgeGroup.Eq(req.AgeGroup))
	}

	classes, err := uc.repo.Class.Where(filter...).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).
		Find()
	if err != nil {
		logger.Error(op, " list class error: ", err)
		return nil, errs.E(op, errs.Database, "list class error")
	}

	return toListClassesResponse(classes), nil
}

func (uc *useCase) DeleteClass(ctx context.Context, req *DeleteClassRequest) (*DeleteClassResponse, error) {
	const op errs.Op = "class.useCase.DeleteClass"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	_, err := uc.repo.Class.Where(uc.repo.Class.ID.Eq(req.ID)).Delete()
	if err != nil {
		logger.Error(op, " delete class error: ", err)
		return nil, errs.E(op, errs.Database, "delete class error")
	}
	return nil, nil
}
