//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package registry

import (
	"context"

	"github.com/google/wire"
)

func InjectedSessionUseCase(
	ctx context.Context,
	provider *Provider,
) session.UseCase {
	wire.Build(BuilderSet, session.NewUseCase)
	return nil
}
