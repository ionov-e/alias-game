package database

import (
	"context"
	"go_telegram_start/internal/database/types"
	"go_telegram_start/pkg/telegram"
)

type DB interface {
	GetLastUpdateID(ctx context.Context) (uint64, error)
	SetLastUpdateID(ctx context.Context, lastUpdateID uint64) error
	GetOrCreateUserInfoFromTelegramUser(ctx context.Context, user telegram.User) (types.UserInfo, error)
	UpdateUserInfo(ctx context.Context, userInfo types.UserInfo) error
}
