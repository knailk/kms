package auth

import (
	"context"
	"kms/app/external/persistence/database/repository"
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
	return nil, nil
}
