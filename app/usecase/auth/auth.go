package auth

import (
	"context"
	"errors"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/authjwt"
	"kms/pkg/helpers"
	"kms/pkg/logger"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
		return nil, errs.E(op, errs.InvalidRequest, "wrong password")
	}

	tokenPair, err := uc.generateToken(string(user.Role), user.Username)
	if err != nil {
		return nil, errs.E(op, err, "failed to generate token pair")
	}

	return &LoginResponse{
		Username:     user.Username,
		Email:        user.Email,
		Role:         string(user.Role),
		ParentName:   user.ParentName,
		FullName:     user.FullName,
		Gender:       user.Gender,
		PhoneNumber:  user.PhoneNumber,
		BirthDate:    user.BirthDate,
		PictureURL:   user.PictureURL,
		Address:      user.Address,
		Longitude:    user.Longitude,
		Latitude:     user.Latitude,
		CreatedAt:    user.CreatedAt,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
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

func (uc *useCase) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	const op errs.Op = "auth.useCase.Register"

	countUser, err := uc.repo.User.Where(uc.repo.User.Username.Eq(req.Username)).Count()
	if err != nil {
		logger.Error(op, " find user failed: ", err)
		return nil, errs.E(op, err)
	}

	if countUser > 0 {
		return nil, errs.E(op, errs.Exist, "user already existed")
	}

	usersRequested, err := uc.repo.UserRequested.
		Where(
			uc.repo.UserRequested.Username.Eq(req.Username),
			uc.repo.UserRequested.Status.Eq(string(entity.UserRequestedStatusPending)),
		).Count()
	if err != nil {
		logger.Error(op, "count user requested errorL ", err)
		return nil, errs.E(op, err)
	}

	if usersRequested > 0 {
		return nil, errs.E(op, errs.Exist, "user request is pending")
	}

	hashedPassword, err := helpers.GenerateHash(req.Password)
	if err != nil {
		logger.Error(op, " failed to hash password: ", err)
		return nil, errs.E(op, err, "failed to hash password")
	}

	// validate something else

	err = uc.repo.UserRequested.Create(&entity.UserRequested{
		ID:          uuid.New(),
		Username:    req.Username,
		FullName:    req.FullName,
		ParentName:  req.ParentName,
		Password:    hashedPassword,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		BirthDate:   req.BirthDate,
		Gender:      req.Gender,
		ClassID:     req.ClassID,
		Status:      entity.UserRequestedStatusPending,
	})
	if err != nil {
		logger.Error(op, " create user request failed: ", err)
		return nil, errs.E(op, err, "create user request failed")
	}

	return &RegisterResponse{}, nil
}

func (uc *useCase) RegisterConfirm(ctx context.Context, req *RegisterConfirmRequest) (*RegisterConfirmResponse, error) {
	const op errs.Op = "auth.useCase.RegisterConfirm"

	userRequested, err := uc.repo.UserRequested.Where(uc.repo.UserRequested.ID.Eq(req.ID)).First()
	if err != nil {
		logger.Error(op, " find user failed: ", err)
		return nil, errs.E(op, err)
	}

	if userRequested.Status != entity.UserRequestedStatusPending {
		return nil, errs.E(op, errs.InvalidRequest, "user is not pending")
	}

	uc.repo.Transaction(func(tx *repository.Query) error {
		if req.Action == entity.UserRequestedStatusApproved {
			err = tx.User.Create(&entity.User{
				Username:    userRequested.Username,
				Password:    userRequested.Password,
				Role:        entity.UserRoleStudent,
				ParentName:  userRequested.ParentName,
				FullName:    userRequested.FullName,
				Gender:      userRequested.Gender,
				Email:       userRequested.Email,
				BirthDate:   userRequested.BirthDate,
				PhoneNumber: userRequested.PhoneNumber,
				PictureURL:  "https://i.pravatar.cc/300",
			})
			if err != nil {
				logger.Error(op, " create user failed: ", err)
				return errs.E(op, err)
			}
		}

		_, err = tx.UserRequested.Where(uc.repo.UserRequested.ID.Eq(req.ID)).Update(uc.repo.UserRequested.Status, req.Action)
		if err != nil {
			logger.Error(op, " update user requested failed: ", err)
			return errs.E(op, err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &RegisterConfirmResponse{
		ClassID:  userRequested.ClassID,
		Username: userRequested.Username,
		Status:   req.Action,
	}, nil
}

func (uc *useCase) RegisterRequestList(ctx context.Context, req *RegisterListRequest) (*RegisterListResponse, error) {
	const op errs.Op = "auth.useCase.RegisterList"

	var cond []gen.Condition

	if req.Status != "" {
		cond = append(cond, uc.repo.UserRequested.Status.Eq(req.Status))
	}

	if req.ClassID != "" {
		if classID, err := uuid.Parse(req.ClassID); err == nil {
			cond = append(cond, uc.repo.UserRequested.ClassID.Eq(classID))
		}
	}

	query := uc.repo.UserRequested.Where(cond...)

	if req.Limit > 0 {
		query = query.Limit(req.Limit).Offset(req.Offset)
	}

	usersRequested, err := query.Preload(uc.repo.UserRequested.Class.Select(uc.repo.Class.ClassName, uc.repo.Class.ID)).Find()
	if err != nil {
		logger.Error(op, " find user requested failed: ", err)
		return nil, errs.E(op, err)
	}

	var rep []*UserRequestedResponse

	copier.CopyWithOption(&rep, &usersRequested, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	return &RegisterListResponse{
		RegisterList: rep,
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
