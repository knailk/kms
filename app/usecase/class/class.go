package class

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/date"
	"kms/pkg/logger"
	"kms/pkg/time_function"
	"time"

	"github.com/google/uuid"
	"gorm.io/gen"
	"gorm.io/gorm/clause"
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

	classID := uuid.New()
	schedules := make([]entity.Schedule, 0)
	for _, s := range req.Schedules {
		schedules = append(schedules, entity.Schedule{
			ID:       uuid.New(),
			ClassID:  classID,
			FromTime: s.FromTime,
			ToTime:   s.ToTime,
			Action:   s.Action,
			Date:     s.Date,
		})
	}

	err := uc.repo.Class.Create(&entity.Class{
		ID:        classID,
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

	if req.TeacherID != "" {
		filter = append(filter, uc.repo.Class.TeacherID.Eq(req.TeacherID))
	}

	if req.DriverID != "" {
		filter = append(filter, uc.repo.Class.DriverID.Eq(req.DriverID))
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

	classes, err := uc.repo.Class.
		Where(filter...).
		Limit(req.Limit).
		Offset((req.Page - 1) * req.Limit).
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

func (uc *useCase) CheckInOut(ctx context.Context, req *CheckInOutRequest) (*CheckInOutResponse, error) {
	const op errs.Op = "class.useCase.CheckInOut"
	users, err := uc.repo.User.Where(uc.repo.User.Username.In(req.Usernames...)).Find()
	if err != nil {
		logger.Error(op, " get users errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	loc, _ := time_function.LoadLocation(entity.TimeZone)

	date, err := date.FromTime(time.Now(), loc)
	if err != nil {
		logger.Error(op, " get date errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	checkInRequests := make([]*entity.CheckInOut, 0)
	for _, user := range users {
		checkInRequests = append(checkInRequests, &entity.CheckInOut{
			ID:       uuid.New(),
			Username: user.Username,
			Action:   req.Action,
			Date:     date.AsDate(),
			ClassID:  req.ClassID,
		})
	}

	err = uc.repo.CheckInOut.Create(checkInRequests...)
	if err != nil {
		logger.Error(op, " create check in errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &CheckInOutResponse{
		Status: "OK",
	}, nil
}

func (uc *useCase) ListMembersInClass(ctx context.Context, req *ListMembersInClassRequest) (*ListMembersInClassResponse, error) {
	const op errs.Op = "class.useCase.ListUsersInClass"

	rep := make([]*GetUserInClass, 0)
	err := uc.repo.User.Select(
		uc.repo.User.ALL,
		uc.repo.UserClass.Status,
		uc.repo.UserClass.CreatedAt.As("joined_at"),
	).LeftJoin(
		uc.repo.UserClass,
		uc.repo.User.Username.EqCol(uc.repo.UserClass.Username),
	).Where(
		uc.repo.UserClass.ClassID.Eq(req.ClassID),
	).Scan(&rep)
	if err != nil {
		logger.Error(op, " get users errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &ListMembersInClassResponse{
		Users: rep,
	}, nil
}

func (uc *useCase) AddMembersToClass(ctx context.Context, req *AddMembersToClassRequest) (*AddMemberToClassResponse, error) {
	const op errs.Op = "class.useCase.AddMembersToClass"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	var usernames []string
	err := uc.repo.User.Select(uc.repo.User.Username).Where(uc.repo.User.Username.In(req.Usernames...)).Scan(&usernames)
	if err != nil {
		logger.Error(op, " get users errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	usersClass := make([]*entity.UserClass, 0)
	for _, user := range usernames {
		usersClass = append(usersClass, &entity.UserClass{
			Username: user,
			ClassID:  req.ClassID,
			Status:   string(entity.UserClassStatusStudying),
		})
	}

	err = uc.repo.UserClass.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}, {Name: "class_id"}},
			DoNothing: true,
		},
	).Create(usersClass...)
	if err != nil {
		logger.Error(op, " create user class errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &AddMemberToClassResponse{}, nil
}

func (uc *useCase) RemoveMembersFromClass(ctx context.Context, req *RemoveMembersFromClassRequest) (*RemoveMembersFromClassResponse, error) {
	const op errs.Op = "class.useCase.RemoveMembersFromClass"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	_, err := uc.repo.UserClass.
		Where(
			uc.repo.UserClass.Username.In(req.Usernames...),
			uc.repo.UserClass.ClassID.Eq(req.ClassID),
		).
		Update(uc.repo.UserClass.Status, entity.UserClassStatusCanceled)
	if err != nil {
		logger.Error(op, " delete user class errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &RemoveMembersFromClassResponse{}, nil
}

func (uc *useCase) CheckInOutHistories(ctx context.Context, req *CheckInOutHistoriesRequest) (*CheckInOutHistoriesResponse, error) {
	const op errs.Op = "class.useCase.CheckInOutHistories"

	if errKind := req.Validate(); errKind != errs.Other {
		return nil, errs.E(op, errKind, "Validate request error")
	}

	checkInOuts, err := uc.repo.CheckInOut.Where(
		uc.repo.CheckInOut.ClassID.Eq(req.ClassID),
		uc.repo.CheckInOut.Date.Eq(req.Date),
	).Find()
	if err != nil {
		logger.Error(op, " get check in out errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &CheckInOutHistoriesResponse{
		Histories: toCheckInOutHistoriesResponse(checkInOuts),
	}, nil
}
