//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"context"

	"github.com/google/wire"
	"kms/app/usecase/auth"
)

func InjectedAuthUseCase(
	ctx context.Context,
	provider *Provider,
) auth.IUseCase {
	wire.Build(BuilderSet, auth.NewUseCase)
	return nil
}
