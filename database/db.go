package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_telegram_start/database/types"
	"go_telegram_start/telegram"
	"strconv"
	"time"
)

type DB struct {
	rc *RedisClient
}

func NewLocal() *DB {
	return &DB{NewClientLocal()}
}

func (db *DB) GetOrCreateUserInfo(ctx context.Context, user telegram.User) (types.UserInfo, error) {
	var userInfo types.UserInfo
	key := strconv.FormatInt(user.ID, 10)
	err := db.rc.Get(ctx, key).Scan(&userInfo)

	if !errors.Is(err, redis.Nil) {
		return userInfo, nil
	}

	newUserInfo := types.UserInfo{
		TelegramID:         user.ID,
		Name:               user.FirstName,
		LastRequestTime:    time.Now(),
		PreferenceLanguage: user.LanguageWithDefault(),
	}
	err = db.rc.Set(ctx, key, newUserInfo, 0).Err()

	if err != nil {
		return newUserInfo, fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return newUserInfo, nil
}

func (db *DB) UpdateUserInfo(ctx context.Context, userInfo types.UserInfo) error {
	key := strconv.FormatInt(userInfo.TelegramID, 10)

	data, err := userInfo.MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshal UserInfo failed: %w", err)
	}

	err = db.rc.Set(ctx, key, data, 0).Err()
	if err != nil {
		return fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return nil
}
