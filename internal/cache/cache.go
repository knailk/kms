package cache

import (
	"context"
	"fmt"
	"kms/app/config"
	"kms/app/external/infra/redis"
	"kms/internal/shutdown"
)

type taskRedis struct {
	client redis.RedisClient
}

func InitRedis(ctx context.Context, tasks *shutdown.Tasks, cfg *config.Config) (redis.RedisClient, error) {
	if tasks.HasStopSignal() {
		return nil, shutdown.ErrAbortedAsGotStopSignal
	}

	client, err := redis.Open(ctx, &cfg.Cache)
	if err != nil {
		return nil, fmt.Errorf("redis open connection: %w", err)
	}

	tasks.Add(&taskRedis{client: client})
	return client, nil
}

func (t *taskRedis) Name() string {
	return "cache"
}

func (t *taskRedis) Shutdown(ctx context.Context) error {
	return t.client.Close()
}
