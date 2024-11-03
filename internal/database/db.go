package database

import (
	"context"
	"go_telegram_start/internal/database/types"
	"go_telegram_start/pkg/telegram"
)

type DB interface {
	GetOrCreateUserInfoFromTelegramUser(ctx context.Context, user telegram.User) (types.UserInfo, error)
	UpdateUserInfo(ctx context.Context, userInfo types.UserInfo) error
}
