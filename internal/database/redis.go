package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_telegram_start/internal/database/types"
	"go_telegram_start/pkg/telegram"
	"log"
	"strconv"
	"time"
)

const lastUpdateIDKey = "lastUpdateID"

type Redis struct {
	rc *redis.Client
}

func NewLocalRedis() *Redis {
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

func (r *Redis) UserInfoFromTelegramUser(ctx context.Context, user telegram.User) (types.UserInfo, error) {
	var userInfo types.UserInfo
	key := strconv.FormatInt(user.ID, 10)
	err := r.rc.Get(ctx, key).Scan(&userInfo)

	if !errors.Is(err, redis.Nil) {
		return userInfo, nil
	}

	newUserInfo := types.UserInfo{
		TelegramID:         user.ID,
		Name:               user.FirstName,
		LastRequestTime:    time.Now(),
		PreferenceLanguage: user.LanguageWithDefault(),
	}
	err = r.rc.Set(ctx, key, newUserInfo, 0).Err()

	if err != nil {
		return newUserInfo, fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return newUserInfo, nil
}

func (r *Redis) SaveUserInfo(ctx context.Context, userInfo types.UserInfo) error {
	key := strconv.FormatInt(userInfo.TelegramID, 10)

	data, err := userInfo.MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshal UserInfo failed: %w", err)
	}

	err = r.rc.Set(ctx, key, data, 0).Err()
	if err != nil {
		return fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return nil
}
