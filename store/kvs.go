package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sugimoto-ne/go_sample_app.git/config"
	"github.com/sugimoto-ne/go_sample_app.git/entity"
)

type KVS struct {
	Cli *redis.Client
}

func NewKVS(ctx context.Context, cfg *config.Config) (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &KVS{Cli: cli}, nil
}

func (k *KVS) Save(ctx context.Context, key string, userID entity.UserID) error {
	id := int64(userID)
	return k.Cli.Set(ctx, key, id, 30*time.Minute).Err()
}

func (k *KVS) Load(ctx context.Context, key string) (entity.UserID, error) {
	id, err := k.Cli.Get(ctx, key).Int64()

	if err != nil {
		return 0, fmt.Errorf("failed to get by %q: %w", key, err)
	}

	return entity.UserID(id), nil
}
