package cache

import (
	"context"
	"fmt"
	"kms/app/external/infra/ristretto"
	"kms/internal/shutdown"
)

type taskRistretto struct {
	client ristretto.RistrettoCache
}

func InitRistretto(ctx context.Context, tasks *shutdown.Tasks) (ristretto.RistrettoCache, error) {
	if tasks.HasStopSignal() {
		return nil, shutdown.ErrAbortedAsGotStopSignal
	}

	client, err := ristretto.NewRistrettoCache()
	if err != nil {
		return nil, fmt.Errorf("redis open connection: %w", err)
	}

	tasks.Add(&taskRistretto{client: client})
	return client, nil
}

func (t *taskRistretto) Name() string {
	return "cache"
}

func (t *taskRistretto) Shutdown(ctx context.Context) error {
	return nil
}
