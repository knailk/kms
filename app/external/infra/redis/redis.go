package redis

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"kms/app/config"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	GetAndScan(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) (int64, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	Close() error
}

type redisClient struct {
	client *redis.Client
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("Not Found", key)
	}
	// nolint: wrapcheck
	return value, err
}

func (r *redisClient) GetAndScan(ctx context.Context, key string, value interface{}) error {
	err := r.client.Get(ctx, key).Scan(value)
	if errors.Is(err, redis.Nil) {
		return fmt.Errorf("Not Found", key)
	}
	// nolint: wrapcheck
	return err
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// nolint: wrapcheck
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisClient) Delete(ctx context.Context, keys ...string) (int64, error) {
	// nolint: wrapcheck
	return r.client.Del(ctx, keys...).Result()
}

func (r *redisClient) Keys(ctx context.Context, pattern string) ([]string, error) {
	// nolint: wrapcheck
	return r.client.Keys(ctx, pattern).Result()
}

func (r *redisClient) Close() error {
	// nolint: wrapcheck
	return r.client.Close()
}

func Open(ctx context.Context, cfg *config.CacheCnf) (RedisClient, error) {
	options, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("redis parse URL: %w", err)
	}

	options.PoolSize = cfg.PoolSize
	if cfg.IdleTimeout > 0 {
		options.IdleTimeout = time.Duration(cfg.IdleTimeout) * time.Second
	}
	if cfg.ReadTimeout > 0 {
		options.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Second
	}
	if cfg.WriteTimeout > 0 {
		options.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Second
	}
	if cfg.MinIdleConns > 0 {
		options.MinIdleConns = cfg.MinIdleConns
	}
	if cfg.UseTLS {
		options.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	client := redis.NewClient(options)
	if err = client.Ping(ctx).Err(); err != nil {
		err = fmt.Errorf("ping redis: %w", err)
		return nil, err
	}

	return &redisClient{client: client}, nil
}

func NewRedisMock() (RedisClient, redismock.ClientMock) {
	client, redisMock := redismock.NewClientMock()
	return &redisClient{client: client}, redisMock
}
