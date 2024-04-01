package check_in_out

import (
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
