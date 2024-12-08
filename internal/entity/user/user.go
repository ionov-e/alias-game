package user

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
	userDB "alias-game/internal/entity/user/db"
	dictionaryEntity "alias-game/internal/entity/user/subentity/dictionary"
	"alias-game/internal/storage"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type dictionaryInterface interface {
	List(ctx context.Context) ([]string, error)
	Word(ctx context.Context, number uint16) (string, error)
}

type User struct {
	userInfo        userDB.UserInfo
	dbForUser       storage.UserDBInterface
	dbForDictionary storage.DictionaryDBInterface
}

func NewFromTelegramUser(ctx context.Context, dbForUser storage.UserDBInterface, dbForDictionary storage.DictionaryDBInterface, tgUser *tgTypes.User) (User, error) {
	userInfo, err := dbForUser.UserInfoFromTelegramUser(ctx, *tgUser)
	if err != nil {
		return User{}, fmt.Errorf("error getting userInfo: %w", err)
	}
	return User{userInfo: userInfo, dbForUser: dbForUser, dbForDictionary: dbForDictionary}, nil
}

func (u *User) TelegramID() int64 {
	return u.userInfo.TelegramID
}

func (u *User) CurrentMenuKey() string {
	return u.userInfo.CurrentMenu
}

func (u *User) ChangeCurrentMenu(ctx context.Context, menuKey menuConstant.Key) error {
	newMenuKey := string(menuKey)
	u.userInfo.CurrentMenu = newMenuKey
	err := u.dbForUser.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("failed ChangeCurrentMenu for user %d with menuConstant %s: %w", u.userInfo.TelegramID, newMenuKey, err)
	}
	return nil
}

func (u *User) Word(ctx context.Context, wordNumber uint16) (string, error) {
	dict, err := u.dictionaryFactoryMethod(u.userInfo.RoundDictionaryKeyAndTry)
	if err != nil {
		return "", fmt.Errorf("RedisDictionaryFactoryMethod failed: %w", err)
	}

	word, err := dict.Word(ctx, wordNumber)
	if err != nil {
		return "", fmt.Errorf("failed at getting Word: %w", err)
	}

	return word, nil
}

func (u *User) UpdateWordResult(ctx context.Context, wordNumber uint16, wordResult userConstant.WordResult) error {
	u.userInfo.AddWordResult(wordNumber, wordResult)
	err := u.dbForUser.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("failed to save updated user info: %w", err)
	}
	return nil
}

func (u *User) ResultStringForTelegram(ctx context.Context) (string, error) {
	dictionaryWords, err := u.dbForDictionary.DictionaryWordList(ctx, u.userInfo.RoundDictionaryKeyAndTry)
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

func (u *User) ChooseDictionary(ctx context.Context, keyForDictionary dictionaryConstant.Key) error {
	u.userInfo.AddLastRequest()

	dictionaryCount, err := u.userInfo.FindDictionaryCountInHistory(keyForDictionary)

	var newDictionaryCount uint16
	if err == nil {
		dictionaryCount.Count++
		newDictionaryCount = dictionaryCount.Count
	} else {
		newDictionaryCount = 1
		u.userInfo.DictionaryHistory = append(u.userInfo.DictionaryHistory, userDB.DictionaryCount{
			DictionaryKey: keyForDictionary,
			Count:         newDictionaryCount,
		})
	}
	u.userInfo.RoundDictionaryKeyAndTry = userDB.NewDictionaryKeyAndTry(keyForDictionary, newDictionaryCount)

	err = u.dbForUser.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("in ChooseDictionary failed to save updated user info: %w", err)
	}

	return nil
}

func (u *User) dictionaryFactoryMethod(keyAndTry userDB.DictionaryKeyAndTry) (dictionaryInterface, error) {
	if keyAndTry.BaseKey == dictionaryConstant.Easy1 {
		return dictionaryEntity.NewEasy1(keyAndTry.TryNumber, u.dbForDictionary), nil
	}

	return nil, fmt.Errorf("unknown dictionary key: %s", keyAndTry.BaseKey)
}
