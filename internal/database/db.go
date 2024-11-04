package database

import (
	dbTypes "alias-game/internal/database/types"
	telegramTypes "alias-game/pkg/telegram/types"
	"context"
)

type DB interface {
	LastUpdateID(ctx context.Context) (uint64, error)
	SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error
	UserInfoFromTelegramUser(ctx context.Context, user telegramTypes.User) (dbTypes.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo dbTypes.UserInfo) error
}
