package user

import (
	"alias-game/internal/database"
	dbConstants "alias-game/internal/database/constants"
	dbTypes "alias-game/internal/database/types"
	"alias-game/internal/service/alias/dictionary"
	telegramTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type User struct {
	userInfo dbTypes.UserInfo
	db       database.DB
}

func UserFromTelegramUser(ctx context.Context, db database.DB, tgUser *telegramTypes.User) (User, error) {
	userInfo, err := db.UserInfoFromTelegramUser(ctx, *tgUser)
	if err != nil {
		return User{}, fmt.Errorf("error getting userInfo: %w", err)
	}
	return User{userInfo: userInfo, db: db}, nil
}

func (u *User) TelegramID() int64 {
	return u.userInfo.TelegramID
}

func (u *User) CurrentMenuKey() string {
	return u.userInfo.CurrentMenu
}

func (u *User) ChangeCurrentMenu(ctx context.Context, menuKey dbConstants.MenuKeyStored) error {
	newMenuKey := menuKey.String()
	u.userInfo.CurrentMenu = newMenuKey
	err := u.db.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("failed ChangeCurrentMenu for user %d with menuKey %s: %w", u.userInfo.TelegramID, newMenuKey, err)
	}
	return nil
}

func (u *User) Word(ctx context.Context, wordNumber uint16) (string, error) {
	dict, err := dictionary.RedisDictionaryFactoryMethod(u.userInfo.RoundDictionaryKeyAndTry, u.db)
	if err != nil {
		return "", fmt.Errorf("RedisDictionaryFactoryMethod failed: %w", err)
	}

	word, err := dict.Word(ctx, wordNumber)
	if err != nil {
		return "", fmt.Errorf("failed at getting Word: %w", err)
	}

	return word, nil
}

func (u *User) UpdateWordResult(ctx context.Context, wordNumber uint16, wordResult dbConstants.WordResult) error {
	u.userInfo.AddWordResult(wordNumber, wordResult)
	err := u.db.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("failed to save updated user info: %w", err)
	}
	return nil
}

func (u *User) ResultStringForTelegram(ctx context.Context) (string, error) {
	dictionaryWords, err := u.db.DictionaryWordList(ctx, u.userInfo.RoundDictionaryKeyAndTry)
	if err != nil {
		return "", fmt.Errorf("failed to save updated user info: %w", err)
	}

	msg := "Round results:\n"
	for number, wordToGuess := range u.userInfo.RoundWords {
		wordAsString := dictionaryWords[wordToGuess.NumberInDictionary]
		msg += fmt.Sprintf("%d) %s %s\n", number, wordAsString, wordToGuess.Result)
	}

	return msg, nil
}

func (u *User) ChooseDictionary(ctx context.Context, dictionaryKey dbConstants.DictionaryKey) error {
	u.userInfo.AddLastRequest()

	dictionaryCount, err := u.userInfo.FindDictionaryCountInHistory(dictionaryKey)

	var newDictionaryCount uint16
	if err == nil {
		dictionaryCount.Count++
		newDictionaryCount = dictionaryCount.Count
	} else {
		newDictionaryCount = 1
		u.userInfo.DictionaryHistory = append(u.userInfo.DictionaryHistory, dbTypes.DictionaryCount{
			DictionaryKey: dictionaryKey,
			Count:         newDictionaryCount,
		})
	}
	u.userInfo.RoundDictionaryKeyAndTry = dbConstants.NewDictionaryKey(dictionaryKey, newDictionaryCount)

	err = u.db.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("in ChooseDictionary failed to save updated user info: %w", err)
	}

	return nil
}
