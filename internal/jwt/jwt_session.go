package jwt

import (
	"context"
	"kms/app/config"
	"kms/internal/shutdown"
	"kms/pkg/authjwt"
	"time"
)

type task struct{}

func InitJWTService(ctx context.Context, tasks *shutdown.Tasks, cfg *config.Config) error {
	if tasks.HasStopSignal() {
		return shutdown.ErrAbortedAsGotStopSignal
	}

	authjwt.InitJWTSession(&authjwt.Config{
		Secret:          cfg.Session.JWT.Secret,
		Issuer:          cfg.Session.JWT.Issuer,
		AccessTokenExp:  time.Duration(cfg.Session.JWT.AccessTokenExp) * time.Second,
		RefreshTokenExp: time.Duration(cfg.Session.JWT.RefreshTokenExp) * time.Second,
	})

	// Add shutdown task
	tasks.Add(&task{})
	return nil
}

func (t *task) Name() string {
	return "session-jwt"
}

func (t *task) Shutdown(ctx context.Context) error {
	return nil
}
