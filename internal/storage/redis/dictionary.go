package redis

import (
	userDB "alias-game/internal/entity/user/db"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"time"                         //nolint:nolintlint,goimports
)

const ttlForDictionary = 365 * 24 * time.Hour

type Dictionary struct {
	client *redis.Client
}

func NewDictionary(client *redis.Client) *Dictionary {
	return &Dictionary{client: client}
}

func (r *Dictionary) DictionaryCreate(ctx context.Context, key userDB.DictionaryKeyAndTry, words []string) error {
	for _, word := range words {
		if err := r.client.RPush(ctx, key.String(), word).Err(); err != nil {
			return fmt.Errorf("adding word %q to key %s in redis failed: %w", word, key, err)
		}
	}
	if err := r.client.Expire(ctx, key.String(), ttlForDictionary).Err(); err != nil {
		return fmt.Errorf("setting expiration for key %s in redis failed: %w", key, err)
	}
	return nil
}

func (r *Dictionary) DictionaryExists(ctx context.Context, key userDB.DictionaryKeyAndTry) (bool, error) {
	exists, err := r.client.Exists(ctx, key.String()).Result()
	if err != nil {
		return false, fmt.Errorf("checking if key %s exists in redis failed: %w", key, err)
	}
	return exists == 1, nil
}

func (r *Dictionary) DictionaryWordList(ctx context.Context, key userDB.DictionaryKeyAndTry) ([]string, error) {
	exists, err := r.client.Exists(ctx, key.String()).Result()
	if err != nil {
		return nil, fmt.Errorf("error checking existence of key %s: %w", key, err)
	}
	if exists == 0 {
		return nil, fmt.Errorf("no dictionary for key %s in redis", key)
	}

	words, err := r.client.LRange(ctx, key.String(), 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("getting dictionary with key %s failed: %w", key, err)
	}
	return words, nil
}

func (r *Dictionary) DictionaryWord(ctx context.Context, key userDB.DictionaryKeyAndTry, index uint16) (string, error) {
	exists, err := r.client.Exists(ctx, key.String()).Result()
	if err != nil {
		return "", fmt.Errorf("error checking existence of key %s: %w", key, err)
	}
	if exists == 0 {
		return "", fmt.Errorf("no dictionary for key %s in redis", key)
	}

	word, err := r.client.LIndex(ctx, key.String(), int64(index)).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("no word at index %d for key %s in redis", index, key)
	}
	if err != nil {
		return "", fmt.Errorf("retrieving word at index %d for key %s failed: %w", index, key, err)
	}
	return word, nil
}
