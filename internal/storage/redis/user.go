package redis

import (
	menuConstant "alias-game/internal/constant/menu"
	userDB "alias-game/internal/entity/user/db"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9" //nolint:nolintlint,goimports
	"strconv"
	"time" //nolint:nolintlint,goimports
)

type User struct {
	client *redis.Client
}

func NewUser(client *redis.Client) *User {
	return &User{client: client}
}

func (r *User) UserInfoFromTelegramUser(ctx context.Context, user tgTypes.User) (userDB.UserInfo, error) {
	var userInfo userDB.UserInfo
	key := r.keyForUserInfo(user.ID)
	err := r.client.Get(ctx, key).Scan(&userInfo)

	if !errors.Is(err, redis.Nil) {
		return userInfo, nil
	}

	newUserInfo := userDB.UserInfo{
		TelegramID:         user.ID,
		Name:               user.FirstName,
		CurrentMenu:        string(menuConstant.Start0Key),
		LastRequestTime:    time.Now(),
		PreferenceLanguage: user.LanguageWithDefault(),
	}
	err = r.client.Set(ctx, key, newUserInfo, 0).Err()

	if err != nil {
		return newUserInfo, fmt.Errorf("setting key %s in redis for creating userInfo failed: %w", key, err)
	}

	return newUserInfo, nil
}

func (r *User) SaveUserInfo(ctx context.Context, userInfo *userDB.UserInfo) error {
	key := r.keyForUserInfo(userInfo.TelegramID)

	err := r.client.Set(ctx, key, userInfo, 0).Err()
	if err != nil {
		return fmt.Errorf("setting key %s in redis for updating userInfo failed: %w", key, err)
	}

	return nil
}

func (r *User) keyForUserInfo(userID int64) string {
	return "user:" + strconv.FormatInt(userID, 10)
}
