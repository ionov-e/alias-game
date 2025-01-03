package user

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
	userDB "alias-game/internal/entity/user/db"
	dictionaryCollection "alias-game/internal/entity/user/dictionary"
	"alias-game/internal/storage"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"math/rand"
	"time"
)

const gameResultString = "Результат игры"
const correctAnswersCountString = "Количество правильных ответов"

type User struct {
	userInfo *userDB.UserInfo
	db       storage.UserDBInterface
}

func NewFromTelegramUser(ctx context.Context, db storage.UserDBInterface, tgUser *tgTypes.User) (*User, error) {
	userInfo, err := db.UserInfoFromTelegramUser(ctx, tgUser)
	if err != nil {
		return nil, fmt.Errorf("error getting userInfo: %w", err)
	}
	return &User{userInfo: userInfo, db: db}, nil
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
	err := u.db.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("failed ChangeCurrentMenu for user %d with menuConstant %s: %w", u.userInfo.TelegramID, newMenuKey, err)
	}
	return nil
}

func (u *User) Word(wordNumber uint16) (string, error) {
	if int(wordNumber) >= len(u.userInfo.RoundWords)-1 {
		return "", fmt.Errorf("wordNumber %d is too much for RoundWords slice of %d", wordNumber, len(u.userInfo.RoundWords))
	}
	return u.userInfo.RoundWords[wordNumber], nil
}

func (u *User) UpdateWordResult(ctx context.Context, wordNumber uint16, wordResult userConstant.WordResult) error {
	u.userInfo.AddWordResult(wordNumber, wordResult)
	err := u.db.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("failed to save updated user info: %w", err)
	}
	return nil
}

func (u *User) EndRound() (string, error) {
	result, err := u.resultString()
	if err != nil {
		return "", fmt.Errorf("in EndRound failed resultString: %w", err)
	}
	u.userInfo.RoundWordResults = []userConstant.WordResult{}
	u.userInfo.RoundWords = []string{}
	u.userInfo.AddLastRequest()
	return result, nil
}

func (u *User) resultString() (string, error) {
	correctAnswers := 0
	roundWords := u.userInfo.RoundWords
	msg := fmt.Sprintf("%s:\n", gameResultString)

	for number, wordResult := range u.userInfo.RoundWordResults {
		msg += fmt.Sprintf("%d) %s %s\n", number+1, roundWords[number], wordResult)
		if wordResult == userConstant.Correct {
			correctAnswers++
		}
	}
	msg += fmt.Sprintf("%s: %d", correctAnswersCountString, correctAnswers)

	return msg, nil
}

func (u *User) ChooseDictionary(ctx context.Context, keyForDictionary dictionaryConstant.Key) error {
	u.userInfo.AddLastRequest()
	u.userInfo.ChooseAnotherDictionary(keyForDictionary)

	wordsFromDict, err := u.wordsFromDictionary()
	if err != nil {
		return fmt.Errorf("wordsFromDictionary failed: %w", err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // We shouldn't bother
	r.Shuffle(len(wordsFromDict), func(i, j int) { wordsFromDict[i], wordsFromDict[j] = wordsFromDict[j], wordsFromDict[i] })
	u.userInfo.RoundWords = wordsFromDict

	err = u.db.SaveUserInfo(ctx, &u.userInfo)
	if err != nil {
		return fmt.Errorf("in ChooseDictionary failed to save updated user info: %w", err)
	}

	return nil
}

func (u *User) wordsFromDictionary() ([]string, error) {
	if u.userInfo.RoundDictionaryKey == dictionaryConstant.Easy1 {
		return dictionaryCollection.Ease1List(), nil
	}

	return nil, fmt.Errorf("unknown dictionary key: %s", u.userInfo.RoundDictionaryKey)
}
