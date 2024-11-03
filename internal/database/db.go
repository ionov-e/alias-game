package database

import (
	"context"
	"go_telegram_start/internal/database/types"
	"go_telegram_start/pkg/telegram"
)

type DB interface {
	LastUpdateID(ctx context.Context) (uint64, error)
	SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error
	UserInfoFromTelegramUser(ctx context.Context, user telegram.User) (types.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo types.UserInfo) error
}
