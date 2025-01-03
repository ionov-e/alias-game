package storage

import (
	userDB "alias-game/internal/entity/user/db"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
)

type UserDBInterface interface {
	UserInfoFromTelegramUser(ctx context.Context, user *tgTypes.User) (*userDB.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo *userDB.UserInfo) error
}
