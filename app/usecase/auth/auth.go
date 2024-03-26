package auth

import (
	"context"
	"errors"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/authjwt"
	"kms/pkg/helpers"
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

func (uc *useCase) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	const op errs.Op = "auth.useCase.Login"

	user, err := uc.repo.User.Where(uc.repo.User.Username.Eq(req.Username)).First()
	if err != nil {
		return nil, errs.E(op, err)
	}

	if user.IsDeleted {
		return nil, errs.E(op, errs.NotExist, "user is deleted")
	}

	if !helpers.ValidateHash(req.Password, user.Password) {
		return nil, errs.E(op, errs.Invalid, "wrong password")
	}

	tokenPair, err := uc.generateToken(string(user.Role), user.Username)
	if err != nil {
		return nil, errs.E(op, err, "failed to generate token pair")
	}

	return &LoginResponse{
		Username:     user.Username,
		Email:        user.Email,
		Role:         string(user.Role),
		FullName:     user.FullName,
		Gender:       user.Gender,
		PhoneNumber:  user.PhoneNumber,
		BirthDate:    user.BirthDate,
		PictureURL:   user.PictureURL,
		Address:      user.Address,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (uc *useCase) GetInfo(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error) {
	const op errs.Op = "auth.useCase.GetInfo"

	user, err := uc.repo.User.Where(uc.repo.User.Username.Eq(req.Username)).First()
	if err != nil {
		return nil, errs.E(op, err)
	}

	return &GetInfoResponse{
		Username:    user.Username,
		Email:       user.Email,
		Role:        string(user.Role),
		FullName:    user.FullName,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		BirthDate:   user.BirthDate,
		PictureURL:  user.PictureURL,
		Address:     user.Address,
	}, nil
}

func (uc *useCase) Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error) {
	const op errs.Op = "auth.useCase.Refresh"
	claims, err := uc.verifyJWTToken(req.RefreshToken)
	if err != nil {
		logger.Error("verify token failed: ", err)
		return nil, errs.E(op, err)
	}

	existedUser, err := uc.repo.User.Where(uc.repo.User.Username.Eq(claims.Username)).Count()
	if err != nil {
		logger.Error("find user failed: ", err)
		return nil, err
	}
	if existedUser == 0 {
		return nil, errs.E(op, errs.NotExist, "user not found")
	}

	tokenPair, err := uc.generateToken(claims.Role, claims.Username)
	if err != nil {
		return nil, errs.E(op, err, "failed to generate token pair")
	}

	return &RefreshResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (uc *useCase) generateToken(
	userType string, username string) (tokenPair *authjwt.TokenPair, err error) {
	uID := uuid.New()

	tokenPair, err = uc.generateJWTTokenPair(uID, userType, username)
	if err != nil {
		return
	}

	return
}

func (uc *useCase) generateJWTTokenPair(
	uID uuid.UUID, userType string, username string,
) (*authjwt.TokenPair, error) {
	claims := authjwt.AuthClaims{
		UID:      uID,
		Role:     userType,
		Username: username,
	}

	tokenPair, err := authjwt.GenerateTokenPair(&claims)
	if err != nil {
		return nil, errors.New("failed to generate token pair")
	}

	return tokenPair, nil
}

func (uc *useCase) verifyJWTToken(token string) (*authjwt.AuthClaims, error) {
	claims, err := authjwt.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
