package user

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/helpers"
	"kms/pkg/logger"
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

	likeKeyword := "%" + req.Keyword + "%"
	users, err := uc.repo.User.
		Where(uc.repo.User.Username.Like(likeKeyword)).
		Or(uc.repo.User.FullName.Like(likeKeyword)).
		Or(uc.repo.User.Email.Like(likeKeyword)).
		Limit(10).
		Find()
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &SearchUserResponse{
		Users: toUsers(users),
	}, nil
}
