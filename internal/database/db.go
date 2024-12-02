package database

import (
	dbConstants "alias-game/internal/database/constants"
	dbTypes "alias-game/internal/database/types"
	telegramTypes "alias-game/pkg/telegram/types"
	"context"
)

type DB interface {
	LastUpdateID(ctx context.Context) (uint64, error)
	SaveLastUpdateID(ctx context.Context, lastUpdateID uint64) error
	UserInfoFromTelegramUser(ctx context.Context, user telegramTypes.User) (dbTypes.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo *dbTypes.UserInfo) error
	DictionaryCreate(ctx context.Context, key dbConstants.DictionaryKeyAndTry, words []string) error
	DictionaryExists(ctx context.Context, key dbConstants.DictionaryKeyAndTry) (bool, error)
	DictionaryWordList(ctx context.Context, key dbConstants.DictionaryKeyAndTry) ([]string, error)
	DictionaryWord(ctx context.Context, key dbConstants.DictionaryKeyAndTry, index uint16) (string, error)
}
