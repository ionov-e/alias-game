package last_update_id

import (
	"context"
	"errors"
	"fmt"                          //nolint:goimports,goimports
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"log"                          //nolint:nolintlint,goimports
	"time"
)

const lastUpdateIDKey = "last-update-id"

type LastUpdateIDRedisClient struct {
	client *redis.Client
}

func NewLastUpdateIDRedisClient(client *redis.Client) LastUpdateIDRedisClient {
	return LastUpdateIDRedisClient{client: client}
}

func (r *LastUpdateIDRedisClient) LastUpdateID(ctx context.Context) (uint64, error) {
	lastUpdateID, err := r.client.Get(ctx, lastUpdateIDKey).Uint64()

	if errors.Is(err, redis.Nil) {
		log.Println("No lastUpdateID in redis")
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
