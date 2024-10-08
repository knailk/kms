package user

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/helpers"
	"kms/pkg/logger"

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

func (uc *useCase) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	const op errs.Op = "auth.useCase.GetUser"

	user, err := uc.repo.User.Where(uc.repo.User.Username.Eq(req.Username)).First()
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &GetUserResponse{
		Username:    user.Username,
		Email:       user.Email,
		Role:        string(user.Role),
		FullName:    user.FullName,
		ParentName:  user.ParentName,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		BirthDate:   user.BirthDate,
		PictureURL:  user.PictureURL,
		Address:     user.Address,
		Longitude:   user.Longitude,
		Latitude:    user.Latitude,
		CreatedAt:   user.CreatedAt,
	}, nil
}

func (uc *useCase) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	const op errs.Op = "auth.useCase.UpdateUser"

	kindErr := req.Validate()
	if kindErr != errs.Other {
		return nil, errs.E(op, kindErr, "validate error")
	}

	if req.Password != "" {
		user, err := uc.repo.User.Where(uc.repo.User.Username.Eq(req.Username)).First()
		if err != nil {
			return nil, errs.E(op, err)
		}

		if !helpers.ValidateHash(req.OldPassword, user.Password) {
			return nil, errs.E(op, errs.InvalidRequest, "wrong password")
		}

		req.Password, err = helpers.GenerateHash(req.Password)
		if err != nil {
			return nil, errs.E(op, errs.InvalidRequest, "generate hash error")
		}
	}

	_, err := uc.repo.User.Where(uc.repo.User.Username.Eq(req.Username)).Updates(&entity.User{
		Password:    req.Password,
		FullName:    req.FullName,
		ParentName:  req.ParentName,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		BirthDate:   req.BirthDate,
		PictureURL:  req.PictureURL,
		Address:     req.Address,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		IsDeleted:   req.IsDeleted,
	})
	if err != nil {
		logger.Error("verify token failed: ", err)
		return nil, errs.E(op, err)
	}

	return &UpdateUserResponse{}, nil
}

func (uc *useCase) SearchUser(ctx context.Context, req *SearchUserRequest) (*SearchUserResponse, error) {
	const op errs.Op = "auth.useCase.SearchUser"

	errKind := req.Validate()
	if errKind != errs.Other {
		logger.Error("validate request failed")
		return nil, errs.E(op, errKind, "validate request failed")
	}

	u := uc.repo.User
	var cond gen.Condition

	if len(req.Roles) > 0 {
		cond = gen.Condition(u.Role.In(req.Roles...))
	}

	likeKeyword := "%" + req.Keyword + "%"
	query := u.
		Where(
			u.Where(cond).Where(
				u.Where(u.Username.Like(likeKeyword)).
					Or(u.FullName.Like(likeKeyword)).
					Or(u.Email.Like(likeKeyword))),
		).Limit(10)

	users, err := query.Find()
	if err != nil {
		logger.Error(op, " search user failed: ", err)
		return nil, errs.E(op, err)
	}

	return &SearchUserResponse{
		Users: toUsers(users),
	}, nil
}

func (uc *useCase) ListTeachersAvailable(ctx context.Context, req *ListTeachersAvailableRequest) (*ListTeachersAvailableResponse, error) {
	const op errs.Op = "auth.useCase.ListTeacherAvailable"

	users, err := uc.repo.User.
		LeftJoin(
			uc.repo.Class,
			uc.repo.Class.TeacherID.EqCol(uc.repo.User.Username),
		).Where(
		uc.repo.User.Role.Eq(entity.UserTypeTeacher),
		uc.repo.Class.ID.IsNull(),
	).Find()
	if err != nil {
		logger.Error(op, " list teacher available failed: ", err)
		return nil, errs.E(op, err)
	}

	return &ListTeachersAvailableResponse{
		Teachers: toUsers(users),
	}, nil
}

func (uc *useCase) ListDriversAvailable(ctx context.Context, req *ListDriversAvailableRequest) (*ListDriversAvailableResponse, error) {
	const op errs.Op = "auth.useCase.ListDriversAvailable"

	users, err := uc.repo.User.
		LeftJoin(
			uc.repo.Class,
			uc.repo.Class.DriverID.EqCol(uc.repo.User.Username),
		).Where(
		uc.repo.User.Role.Eq(entity.UserTypeDriver),
		uc.repo.Class.ID.IsNull(),
	).Find()
	if err != nil {
		logger.Error(op, " list drivers available failed: ", err)
		return nil, errs.E(op, err)
	}

	return &ListDriversAvailableResponse{
		Drivers: toUsers(users),
	}, nil
}
