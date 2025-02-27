package last_update_id

import (
	"context"
	"errors"
	"fmt"                          //nolint:nolintlint,goimports
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"log/slog"
	"time" //nolint:nolintlint,goimports
)

const lastUpdateIDKey = "last-update-id"

type LastUpdateIDRedisClient struct {
	client *redis.Client
	log    *slog.Logger
}

func NewLastUpdateIDRedisClient(client *redis.Client, log *slog.Logger) LastUpdateIDRedisClient {
	return LastUpdateIDRedisClient{client: client, log: log}
}

func (r *LastUpdateIDRedisClient) LastUpdateID(ctx context.Context) (uint64, error) {
	lastUpdateID, err := r.client.Get(ctx, lastUpdateIDKey).Uint64()

	if errors.Is(err, redis.Nil) {
		r.log.Warn("No lastUpdateID in redis")
		return uint64(0), nil
	}

	if err != nil {
		return uint64(0), fmt.Errorf("getting lastUpdateID from redis failed: %w", err)
	}

	return lastUpdateID, nil
}

func (r *LastUpdateIDRedisClient) SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // Ensure the context is canceled to release resources

	if err := r.client.Set(ctx, lastUpdateIDKey, lastUpdateID, 0).Err(); err != nil {
		return fmt.Errorf("setting lastUpdateID failed: %w", err)
	}

	return nil
}
