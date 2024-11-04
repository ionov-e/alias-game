package alias

import (
	"alias-game/internal/database"
	dbTypes "alias-game/internal/database/types"
	telegramTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"time"
)

type User struct {
	userInfo dbTypes.UserInfo
	db       database.DB
}

func NewFromDB(ctx context.Context, db database.DB, tgUser telegramTypes.User) (User, error) {
	userInfo, err := db.UserInfoFromTelegramUser(ctx, tgUser)
	if err != nil {
		return User{}, fmt.Errorf("error getting userInfo: %w", err)
	}
	return User{userInfo: userInfo, db: db}, nil
}

func (u *User) AddNewWord(ctx context.Context, newWord string) error {
	u.userInfo.AddNewWord(newWord)
	u.userInfo.LastRequestTime = time.Now()
	err := u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return fmt.Errorf("error updating userInfo: %w", err)
	}
	return nil
}
