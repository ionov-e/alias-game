package database

import (
	"context"
	dbTypes "go_telegram_start/internal/database/types"
	telegramTypes "go_telegram_start/pkg/telegram/types"
)

type DB interface {
	LastUpdateID(ctx context.Context) (uint64, error)
	SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error
	UserInfoFromTelegramUser(ctx context.Context, user telegramTypes.User) (dbTypes.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo dbTypes.UserInfo) error
}
