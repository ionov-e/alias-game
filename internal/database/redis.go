package database

import (
	dbConstants "alias-game/internal/database/constants"
	dbTypes "alias-game/internal/database/types"
	telegramTypes "alias-game/pkg/telegram/types"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"log"
	"strconv"
	"time" //nolint:goimports
)

const lastUpdateIDKey = "last-update-id"

type Redis struct {
	rc *redis.Client
}

const ttlForDictionary = 365 * 24 * time.Hour

func NewLocalRedis() DB {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default Redis
	}
	return &Redis{
		rc: redis.NewClient(options),
	}
}

func (r *Redis) LastUpdateID(ctx context.Context) (uint64, error) {
	lastUpdateID, err := r.rc.Get(ctx, lastUpdateIDKey).Uint64()

	if errors.Is(err, redis.Nil) {
		log.Println("No lastUpdateID in redis")
		return uint64(0), nil
	}

	if err != nil {
		return uint64(0), fmt.Errorf("getting lastUpdateID from redis failed: %w", err)
	}

	return lastUpdateID, nil
}

func (r *Redis) SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error {
	if err := r.rc.Set(ctx, lastUpdateIDKey, lastUpdateID, 0).Err(); err != nil {
		return fmt.Errorf("setting lastUpdateID failed: %w", err)
	}

	return nil
}

func (r *Redis) UserInfoFromTelegramUser(ctx context.Context, user telegramTypes.User) (dbTypes.UserInfo, error) {
	var userInfo dbTypes.UserInfo
	key := r.keyForUserInfo(user.ID)
	err := r.rc.Get(ctx, key).Scan(&userInfo)

	if !errors.Is(err, redis.Nil) {
		return userInfo, nil
	}

	newUserInfo := dbTypes.UserInfo{
		TelegramID:         user.ID,
		Name:               user.FirstName,
		CurrentMenu:        string(dbConstants.MenuStart0Key),
		LastRequestTime:    time.Now(),
		PreferenceLanguage: user.LanguageWithDefault(),
	}
	err = r.rc.Set(ctx, key, newUserInfo, 0).Err()

	if err != nil {
		return newUserInfo, fmt.Errorf("setting key %s in redis for creating userInfo failed: %w", key, err)
	}

	return newUserInfo, nil
}

func (r *Redis) SaveUserInfo(ctx context.Context, userInfo *dbTypes.UserInfo) error {
	key := r.keyForUserInfo(userInfo.TelegramID)

	err := r.rc.Set(ctx, key, userInfo, 0).Err()
	if err != nil {
		return fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return nil
}

func (r *Redis) keyForUserInfo(userID int64) string {
	return "user:" + strconv.FormatInt(userID, 10)
}

func (r *Redis) DictionaryCreate(ctx context.Context, key dbConstants.DictionaryKeyAndTry, words []string) error {
	for _, word := range words {
		if err := r.rc.RPush(ctx, key.String(), word).Err(); err != nil {
			return fmt.Errorf("adding word %q to key %s in redis failed: %w", word, key, err)
		}
	}
	if err := r.rc.Expire(ctx, key.String(), ttlForDictionary).Err(); err != nil {
		return fmt.Errorf("setting expiration for key %s in redis failed: %w", key, err)
	}
	return nil
}

func (r *Redis) DictionaryExists(ctx context.Context, key dbConstants.DictionaryKeyAndTry) (bool, error) {
	exists, err := r.rc.Exists(ctx, key.String()).Result()
	if err != nil {
		return false, fmt.Errorf("checking if key %s exists in redis failed: %w", key, err)
	}
	return exists == 1, nil
}

func (r *Redis) DictionaryWordList(ctx context.Context, key dbConstants.DictionaryKeyAndTry) ([]string, error) {
	exists, err := r.rc.Exists(ctx, key.String()).Result()
	if err != nil {
		return nil, fmt.Errorf("error checking existence of key %s: %w", key, err)
	}
	if exists == 0 {
		return nil, fmt.Errorf("no dictionary for key %s in redis", key)
	}

	words, err := r.rc.LRange(ctx, key.String(), 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("getting dictionary with key %s failed: %w", key, err)
	}
	return words, nil
}

func (r *Redis) DictionaryWord(ctx context.Context, key dbConstants.DictionaryKeyAndTry, index uint16) (string, error) {
	exists, err := r.rc.Exists(ctx, key.String()).Result()
	if err != nil {
		return "", fmt.Errorf("error checking existence of key %s: %w", key, err)
	}
	if exists == 0 {
		return "", fmt.Errorf("no dictionary for key %s in redis", key)
	}

	word, err := r.rc.LIndex(ctx, key.String(), int64(index)).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("no word at index %d for key %s in redis", index, key)
	}
	if err != nil {
		return "", fmt.Errorf("retrieving word at index %d for key %s failed: %w", index, key, err)
	}
	return word, nil
}
