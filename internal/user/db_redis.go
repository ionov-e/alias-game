package user

import (
	menuConstant "alias-game/internal/constant/menu"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"log/slog"
	"strconv"
	"time" //nolint:nolintlint,goimports
)

// RedisClient implements methods needed to persist user data
type RedisClient struct {
	client *redis.Client
	log    *slog.Logger
}

func NewRedisClient(client *redis.Client, log *slog.Logger) *RedisClient {
	return &RedisClient{client: client, log: log}
}

func (r *RedisClient) userDataFromTelegramUser(ctx context.Context, u *tgTypes.User) (*data, error) {
	var info data
	key := r.keyForUserInfo(u.ID)
	err := r.client.Get(ctx, key).Scan(&info)

	if !errors.Is(err, redis.Nil) {
		return &info, nil
	}

	newUserInfo := data{
		TelegramID:         u.ID,
		Name:               u.FirstName,
		CurrentMenu:        string(menuConstant.Start0),
		LastRequestTime:    time.Now(),
		PreferenceLanguage: u.LanguageWithDefault(),
	}
	err = r.client.Set(ctx, key, newUserInfo, 0).Err()

	if err != nil {
		return &newUserInfo, fmt.Errorf("setting key %s in redis for creating data failed: %w", key, err)
	}

	return &newUserInfo, nil
}

func (r *RedisClient) userDataFromTelegramUserID(ctx context.Context, tgUserID int64) (*data, error) {
	var info data
	key := r.keyForUserInfo(tgUserID)
	err := r.client.Get(ctx, key).Scan(&info)

	if !errors.Is(err, redis.Nil) {
		return &info, nil
	}

	return nil, fmt.Errorf("failed getting user from redis with key %s: %w", tgUserID, err)
}

func (r *RedisClient) saveUserInfo(ctx context.Context, userInfo *data) error {
	key := r.keyForUserInfo(userInfo.TelegramID)

	err := r.client.Set(ctx, key, userInfo, 0).Err()
	if err != nil {
		return fmt.Errorf("setting key %s in redis for updating data failed: %w", key, err)
	}

	return nil
}

func (r *RedisClient) keyForUserInfo(userID int64) string {
	return "user:" + strconv.FormatInt(userID, 10)
}
