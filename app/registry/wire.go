//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package registry 

import (
	"context"

	"github.com/google/wire"
	"kms/app/usecase/auth"
	"kms/app/usecase/user"
	"kms/app/usecase/chat"
	"kms/app/usecase/class"
	"kms/app/usecase/cron"
)

func InjectedAuthUseCase(
	ctx context.Context,
	provider *Provider,
) auth.IUseCase {
	wire.Build(BuilderSet, auth.NewUseCase)
	return nil
}

func InjectedUserUseCase(
	ctx context.Context,
	provider *Provider,
) user.IUseCase {
	wire.Build(BuilderSet, user.NewUseCase)
	return nil
}

func InjectedChatUseCase(
	ctx context.Context,
	provider *Provider,
) chat.IUseCase {
	wire.Build(BuilderSet, chat.NewUseCase)
	return nil
}

func InjectedClassUseCase(
	ctx context.Context,
	provider *Provider,
) class.IUseCase {
	wire.Build(BuilderSet, class.NewUseCase)
	return nil
}

func InjectedCronUseCase(
	ctx context.Context,
	provider *Provider,
) cron.IUseCase {
	wire.Build(BuilderSet, cron.NewUseCase)
	return nil
}


