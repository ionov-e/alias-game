package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"log"                          //nolint:nolintlint,goimports
)

const lastUpdateIDKey = "last-update-id"

type LastUpdateID struct {
	client *redis.Client
}

func NewLastUpdateID(client *redis.Client) LastUpdateID {
	return LastUpdateID{client: client}
}

func (r *LastUpdateID) LastUpdateID(ctx context.Context) (uint64, error) {
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

func (r *LastUpdateID) SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error {
	if err := r.client.Set(ctx, lastUpdateIDKey, lastUpdateID, 0).Err(); err != nil {
		return fmt.Errorf("setting lastUpdateID failed: %w", err)
	}

	return nil
}
