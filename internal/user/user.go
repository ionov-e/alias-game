package user

import (
	dictionaryConstant "alias-game/internal/constant/dictionary"
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
	dictionaryCollection "alias-game/internal/user/dictionary"
	"alias-game/internal/user/vo"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const correctAnswersCountString = "Правильных ответов"
const incorrectAnswersCountString = "Неправильных ответов"
const skippedAnswersCountString = "Пропущенных ответов"

type User struct {
	userInfo *vo.UserInfo
	db       UserDBInterface
}

type UserDBInterface interface {
	UserInfoFromTelegramUser(ctx context.Context, user *tgTypes.User) (*vo.UserInfo, error)
	SaveUserInfo(ctx context.Context, userInfo *vo.UserInfo) error
}

func NewFromTelegramUser(ctx context.Context, db UserDBInterface, tgUser *tgTypes.User) (*User, error) {
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
	u.LastRoundResult()
	newMenuKey := string(menuKey)
	u.userInfo.CurrentMenu = newMenuKey
	err := u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return fmt.Errorf("failed ChangeCurrentMenu for user %d with menuConstant %s: %w", u.userInfo.TelegramID, newMenuKey, err)
	}
	return nil
}

func (u *User) CurrentWord() (string, error) {
	word, err := u.userInfo.Word(u.userInfo.RoundWordNumber)
	if err != nil {
		return "", fmt.Errorf("failed getting current word: %w", err)
	}
	return word, nil
}

func (u *User) SetCurrentWordResult(result userConstant.WordResult) {
	u.setRoundWordResult(u.userInfo.RoundWordNumber, result)
}

func (u *User) setRoundWordResult(wordNumber uint16, wordResult userConstant.WordResult) {
	u.userInfo.SetRoundWordResult(wordNumber, wordResult)
}

func (u *User) ConcludeRound(ctx context.Context) (roundResults string, err error) {
	u.userInfo.AddLastRequest()

	correctAnswers := 0
	incorrectAnswers := 0
	skippedAnswers := 0
	var msg string
	for number, wordResult := range u.userInfo.RoundWordResults {
		msg += fmt.Sprintf("%d) %s %s\n", number+1, wordResult, u.userInfo.RoundWords[number])
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
			return "", fmt.Errorf("wordResult is %v, in user %d", wordResult, u.userInfo.TelegramID)
		}
	}
	msg += fmt.Sprintf("%s: %d\n", correctAnswersCountString, correctAnswers)
	msg += fmt.Sprintf("%s: %d\n", incorrectAnswersCountString, incorrectAnswers)
	msg += fmt.Sprintf("%s: %d\n", skippedAnswersCountString, skippedAnswers)

	err = u.userInfo.AddRoundResult(correctAnswers, incorrectAnswers, skippedAnswers)
	if err != nil {
		return "", fmt.Errorf("in ConcludeRound failed resultString: %w", err)
	}

	err = u.prepareRoundWordsFromDictionary()
	if err != nil {
		return "", err
	}

	teamCount := u.userInfo.AllTeamCount()
	if teamCount > 1 {
		if u.userInfo.RoundTeamNumber == teamCount-1 {
			u.userInfo.RoundTeamNumber = 0
		} else {
			u.userInfo.RoundTeamNumber++
		}
	}

	err = u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return "", fmt.Errorf("in ConcludeRound failed to save updated user info: %w", err)
	}

	return msg, nil
}

func (u *User) AllTeamsCount() uint16 {
	return u.userInfo.AllTeamCount()
}

func (u *User) IsGameEnded() (bool, error) {
	gameEnds, err := u.userInfo.IsGameEnd()
	if err != nil {
		return false, fmt.Errorf("failed userInfo IsGameEnded: %w", err)
	}
	return gameEnds, nil
}

func (u *User) SetRoundTime(ctx context.Context, newRoundTimeInSeconds uint16) error {
	u.userInfo.AddLastRequest()
	u.userInfo.PreferenceRoundTime = newRoundTimeInSeconds
	err := u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return fmt.Errorf("in SetRoundTimePredefined failed to save updated user info: %w", err)
	}
	return nil
}

func (u *User) ChooseDictionary(ctx context.Context, keyForDictionary dictionaryConstant.Key) error {
	u.userInfo.AddLastRequest()
	u.userInfo.ChooseAnotherDictionary(keyForDictionary)

	err := u.prepareRoundWordsFromDictionary()
	if err != nil {
		return err
	}

	err = u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return fmt.Errorf("in ChooseDictionary failed to save updated user info: %w", err)
	}

	return nil
}

func (u *User) prepareRoundWordsFromDictionary() error {
	wordsFromDict, err := u.wordsFromDictionary()
	if err != nil {
		return fmt.Errorf("wordsFromDictionary failed: %w", err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec // We shouldn't bother
	r.Shuffle(len(wordsFromDict), func(i, j int) { wordsFromDict[i], wordsFromDict[j] = wordsFromDict[j], wordsFromDict[i] })
	u.userInfo.RoundWords = wordsFromDict
	u.userInfo.RoundWordResults = []userConstant.WordResult{}
	u.userInfo.RoundWordNumber = 0
	return nil
}

func (u *User) wordsFromDictionary() ([]string, error) {
	if u.userInfo.RoundDictionaryKey == dictionaryConstant.Easy1 {
		return dictionaryCollection.Ease1List(), nil
	}

	return nil, fmt.Errorf("unknown dictionary key: %s", u.userInfo.RoundDictionaryKey)
}

func (u *User) NextWord() {
	u.userInfo.AddLastRequest()
	u.userInfo.RoundWordNumber++
}

func (u *User) SetTeamCount(count uint16) {
	u.userInfo.AddLastRequest()
	u.userInfo.AllTeamsInfo = make([]vo.TeamInfo, count)
}

func (u *User) InfoForFillingTeamNames() (firstTeamNumberWithoutName, totalTeamCount uint16, err error) {
	for i, team := range u.userInfo.AllTeamsInfo {
		if team.Name == "" {
			return uint16(i), uint16(len(u.userInfo.AllTeamsInfo)), nil
		}
	}
	return 0, 0, errors.New("No team without name found in AllTeamsInfo: " + fmt.Sprint(u.userInfo.AllTeamsInfo))
}

// CurrentTeamName returns name of current team or empty if no multiple teams
func (u *User) CurrentTeamName() string {
	if len(u.userInfo.AllTeamsInfo) <= 1 {
		return ""
	}

	return u.userInfo.AllTeamsInfo[u.userInfo.RoundTeamNumber].Name
}

func (u *User) SetTeamName(ctx context.Context, message string, firstTeamNumberWithoutName uint16) error {
	u.userInfo.AddLastRequest()
	u.userInfo.AllTeamsInfo[firstTeamNumberWithoutName].Name = message
	err := u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return fmt.Errorf("in SetTeamName failed to save updated user info: %w", err)
	}
	return nil
}

func (u *User) SetWordCountToWin(wordCountToWin uint16) {
	u.userInfo.AddLastRequest()
	u.userInfo.WordCountToWin = wordCountToWin
}

func (u *User) LastRoundResult() (words []string, results []userConstant.WordResult) {
	u.userInfo.AddLastRequest()
	return u.userInfo.RoundWords, u.userInfo.RoundWordResults
}

func (u *User) GameResult() []userConstant.TeamInfo {
	u.userInfo.AddLastRequest()
	return convertTeamInfo(u.userInfo.AllTeamsInfo)
}

func convertTeamInfo(info []vo.TeamInfo) []userConstant.TeamInfo {
	converted := make([]userConstant.TeamInfo, len(info))

	for i, teamInfo := range info {
		roundResults := make([]userConstant.RoundResult, len(teamInfo.RoundResults))
		var totalCorrect, totalIncorrect, totalSkipped uint16

		for j, rr := range teamInfo.RoundResults {
			roundResults[j] = userConstant.RoundResult{
				CorrectAnswersCount:   rr.CorrectAnswersCount,
				IncorrectAnswersCount: rr.IncorrectAnswersCount,
				SkippedAnswersCount:   rr.SkippedAnswersCount,
			}
			totalCorrect += rr.CorrectAnswersCount
			totalIncorrect += rr.IncorrectAnswersCount
			totalSkipped += rr.SkippedAnswersCount
		}

		converted[i] = userConstant.TeamInfo{
			Name:                       teamInfo.Name,
			RoundResults:               roundResults,
			TotalCorrectAnswersCount:   totalCorrect,
			TotalIncorrectAnswersCount: totalIncorrect,
			TotalSkippedAnswersCount:   totalSkipped,
		}
	}
	return converted
}

func (u *User) ClearGame(ctx context.Context) error {
	u.userInfo.AddLastRequest()
	u.userInfo.RoundWords = []string{}
	u.userInfo.RoundWordResults = []userConstant.WordResult{}
	u.userInfo.RoundWordNumber = 0
	u.userInfo.RoundTeamNumber = 0
	err := u.db.SaveUserInfo(ctx, u.userInfo)
	if err != nil {
		return fmt.Errorf("in ClearGame failed to save updated user info: %w", err)
	}
	return nil
}
