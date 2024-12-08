package storage

import (
	userDB "alias-game/internal/entity/user/db"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
)

type UserDBInterface interface {
	UserInfoFromTelegramUser(ctx context.Context, user tgTypes.User) (userDB.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo *userDB.UserInfo) error
}

type DictionaryDBInterface interface {
	DictionaryWordList(ctx context.Context, key userDB.DictionaryKeyAndTry) ([]string, error)
	DictionaryCreate(ctx context.Context, key userDB.DictionaryKeyAndTry, words []string) error
	DictionaryExists(ctx context.Context, key userDB.DictionaryKeyAndTry) (bool, error)
	DictionaryWord(ctx context.Context, key userDB.DictionaryKeyAndTry, index uint16) (string, error)
}
