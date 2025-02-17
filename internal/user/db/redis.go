package db

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user/vo"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"strconv"
	"time" //nolint:nolintlint,goimports
)

type UserRedisClient struct {
	client *redis.Client
}

func NewUser(client *redis.Client) *UserRedisClient {
	return &UserRedisClient{client: client}
}

func (r *UserRedisClient) UserInfoFromTelegramUser(ctx context.Context, user *tgTypes.User) (*vo.UserInfo, error) {
	var userInfo vo.UserInfo
	key := r.keyForUserInfo(user.ID)
	err := r.client.Get(ctx, key).Scan(&userInfo)

	if !errors.Is(err, redis.Nil) {
		return &userInfo, nil
	}

	newUserInfo := vo.UserInfo{
		TelegramID:         user.ID,
		Name:               user.FirstName,
		CurrentMenu:        string(menuConstant.Start0),
		LastRequestTime:    time.Now(),
		PreferenceLanguage: user.LanguageWithDefault(),
	}
	err = r.client.Set(ctx, key, newUserInfo, 0).Err()

	if err != nil {
		return &newUserInfo, fmt.Errorf("setting key %s in redis for creating userInfo failed: %w", key, err)
	}

	return &newUserInfo, nil
}

func (r *UserRedisClient) SaveUserInfo(ctx context.Context, userInfo *vo.UserInfo) error {
	key := r.keyForUserInfo(userInfo.TelegramID)

	err := r.client.Set(ctx, key, userInfo, 0).Err()
	if err != nil {
		return fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return nil
}

func (r *UserRedisClient) keyForUserInfo(userID int64) string {
	return "user:" + strconv.FormatInt(userID, 10)
}
