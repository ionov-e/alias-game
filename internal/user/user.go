package user

import (
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"errors"
	"fmt"
)

const correctAnswersCountString = "Правильных ответов"
const incorrectAnswersCountString = "Неправильных ответов"
const skippedAnswersCountString = "Пропущенных ответов"

type User struct {
	data *data
	db   DBForUserInterface
}

type DBForUserInterface interface {
	userDataFromTelegramUser(ctx context.Context, user *tgTypes.User) (*data, error)
	saveUserInfo(ctx context.Context, userInfo *data) error
}

func NewFromTelegramUser(ctx context.Context, db DBForUserInterface, tgUser *tgTypes.User) (*User, error) {
	info, err := db.userDataFromTelegramUser(ctx, tgUser)
	if err != nil {
		return nil, fmt.Errorf("error getting data: %w", err)
	}
	return &User{data: info, db: db}, nil
}

func (u *User) TelegramID() int64 {
	return u.data.TelegramID
}

func (u *User) CurrentMenuKey() string {
	return u.data.CurrentMenu
}

func (u *User) ChangeCurrentMenu(ctx context.Context, menuKey menuConstant.Key) error {
	u.LastRoundResult()
	newMenuKey := string(menuKey)
	u.data.CurrentMenu = newMenuKey
	err := u.db.saveUserInfo(ctx, u.data)
	if err != nil {
		return fmt.Errorf("failed ChangeCurrentMenu for user %d with menuConstant %s: %w", u.data.TelegramID, newMenuKey, err)
	}
	return nil
}

func (u *User) CurrentWord() (string, error) {
	word, err := u.data.word(u.data.RoundWordNumber)
	if err != nil {
		return "", fmt.Errorf("failed getting current word: %w", err)
	}
	return word, nil
}

func (u *User) SetCurrentWordResult(result userConstant.WordResult) {
	u.data.setRoundWordResult(u.data.RoundWordNumber, result)
}

func (u *User) ConcludeRound(ctx context.Context) (roundResults string, err error) {
	u.data.addLastRequest()

	correctAnswers := 0
	incorrectAnswers := 0
	skippedAnswers := 0
	var msg string
	for number, wordResult := range u.data.RoundWordResults {
		msg += fmt.Sprintf("%d) %s %s\n", number+1, wordResult, u.data.RoundWords[number])
		switch wordResult {
		case userConstant.Correct:
			correctAnswers++
		case userConstant.Incorrect:
			incorrectAnswers++
		case userConstant.Skipped:
			skippedAnswers++
		case userConstant.NotAnswered:
			continue
		default:
			return "", fmt.Errorf("wordResult is %v, in user %d", wordResult, u.data.TelegramID)
		}
	}
	msg += fmt.Sprintf("%s: %d\n", correctAnswersCountString, correctAnswers)
	msg += fmt.Sprintf("%s: %d\n", incorrectAnswersCountString, incorrectAnswers)
	msg += fmt.Sprintf("%s: %d\n", skippedAnswersCountString, skippedAnswers)

	err = u.data.addRoundResult(correctAnswers, incorrectAnswers, skippedAnswers)
	if err != nil {
		return "", fmt.Errorf("in ConcludeRound failed resultString: %w", err)
	}

	err = u.data.prepareRoundWordsFromDictionary()
	if err != nil {
		return "", err
	}

	teamCount := u.data.allTeamCount()
	if teamCount > 1 {
		if u.data.RoundTeamNumber == teamCount-1 {
			u.data.RoundTeamNumber = 0
		} else {
			u.data.RoundTeamNumber++
		}
	}

	err = u.db.saveUserInfo(ctx, u.data)
	if err != nil {
		return "", fmt.Errorf("in ConcludeRound failed to save updated user data: %w", err)
	}

	return msg, nil
}

func (u *User) AllTeamsCount() uint16 {
	return u.data.allTeamCount()
}

func (u *User) IsGameEnded() (bool, error) {
	gameEnds, err := u.data.isGameEnd()
	if err != nil {
		return false, fmt.Errorf("failed data IsGameEnded: %w", err)
	}
	return gameEnds, nil
}

func (u *User) SetRoundTime(ctx context.Context, newRoundTimeInSeconds uint16) error {
	u.data.addLastRequest()
	u.data.PreferenceRoundTime = newRoundTimeInSeconds
	err := u.db.saveUserInfo(ctx, u.data)
	if err != nil {
		return fmt.Errorf("in SetRoundTimePredefined failed to save updated user data: %w", err)
	}
	return nil
}

func (u *User) ChooseDictionary(ctx context.Context, keyForDictionary userConstant.DictionaryKey) error {
	u.data.addLastRequest()
	u.data.chooseAnotherDictionary(keyForDictionary)

	err := u.data.prepareRoundWordsFromDictionary()
	if err != nil {
		return err
	}

	err = u.db.saveUserInfo(ctx, u.data)
	if err != nil {
		return fmt.Errorf("in ChooseDictionary failed to save updated user data: %w", err)
	}

	return nil
}

func (u *User) NextWord() {
	u.data.addLastRequest()
	u.data.RoundWordNumber++
}

func (u *User) SetTeamCount(count uint16) {
	u.data.addLastRequest()
	u.data.AllTeamsInfo = make([]teamInfo, count)
}

func (u *User) InfoForFillingTeamNames() (firstTeamNumberWithoutName, totalTeamCount uint16, err error) {
	for i, team := range u.data.AllTeamsInfo {
		if team.Name == "" {
			return uint16(i), uint16(len(u.data.AllTeamsInfo)), nil
		}
	}
	return 0, 0, errors.New("No team without name found in AllTeamsInfo: " + fmt.Sprint(u.data.AllTeamsInfo))
}

// CurrentTeamName returns name of current team or empty if no multiple teams
func (u *User) CurrentTeamName() string {
	if len(u.data.AllTeamsInfo) <= 1 {
		return ""
	}

	return u.data.AllTeamsInfo[u.data.RoundTeamNumber].Name
}

func (u *User) SetTeamName(ctx context.Context, message string, firstTeamNumberWithoutName uint16) error {
	u.data.addLastRequest()
	u.data.AllTeamsInfo[firstTeamNumberWithoutName].Name = message
	err := u.db.saveUserInfo(ctx, u.data)
	if err != nil {
		return fmt.Errorf("in SetTeamName failed to save updated user data: %w", err)
	}
	return nil
}

func (u *User) SetWordCountToWin(wordCountToWin uint16) {
	u.data.addLastRequest()
	u.data.WordCountToWin = wordCountToWin
}

func (u *User) LastRoundResult() (words []string, results []userConstant.WordResult) {
	u.data.addLastRequest()
	return u.data.RoundWords, u.data.RoundWordResults
}

func (u *User) GameResult() []userConstant.TeamInfo {
	u.data.addLastRequest()
	return u.data.convertTeamInfo()
}

func (u *User) ClearGame(ctx context.Context) error {
	u.data.addLastRequest()
	u.data.RoundWords = []string{}
	u.data.RoundWordResults = []userConstant.WordResult{}
	u.data.RoundWordNumber = 0
	u.data.RoundTeamNumber = 0
	err := u.db.saveUserInfo(ctx, u.data)
	if err != nil {
		return fmt.Errorf("in ClearGame failed to save updated user data: %w", err)
	}
	return nil
}
