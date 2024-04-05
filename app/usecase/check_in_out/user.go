package check_in_out

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
)

type useCase struct {
	repo *repository.PostgresRepository
}

func NewUseCase(repo *repository.PostgresRepository) IUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) CheckInOut(ctx context.Context, req *CheckInOutRequest) (*CheckInOutResponse, error) {
	const op errs.Op = "check_in_out.useCase.CheckInOut"
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
		})
	}

	err = uc.repo.CheckInOut.Create(checkInRequests...)
	if err != nil {
		logger.Error(op, " create check in errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &CheckInOutResponse{}, nil
}

func (uc *useCase) ListUsersInClass(ctx context.Context, req *ListUsersInClassRequest) (*ListUsersInClassResponse, error) {
	const op errs.Op = "check_in_out.useCase.ListUsersInClass"
	users, err := uc.repo.User.
		LeftJoin(
			uc.repo.UserClass,
			uc.repo.User.Username.EqCol(uc.repo.UserClass.Username),
		).Where(
		uc.repo.UserClass.ClassID.Eq(req.ClassID),
		uc.repo.UserClass.Status.Eq("studying"),
	).Find()
	if err != nil {
		logger.Error(op, " get users errors ", err)
		return nil, errs.E(op, errs.Database, err)
	}

	return &ListUsersInClassResponse{
		Users: toUsersInClass(users),
	}, nil
}
