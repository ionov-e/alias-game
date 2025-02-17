package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	userConstant "alias-game/internal/constant/user"
	userEntity "alias-game/internal/entity/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

const nextInCurrentGameMessage = "Продолжить"
const startAnewMessage = "Прекратить текущую игру"

type CurrentGameResult struct {
	tgClient *telegram.Client
	user     *userEntity.User
}

func NewCurrentGameResult(tgClient *telegram.Client, user *userEntity.User) CurrentGameResult {
	return CurrentGameResult{
		tgClient: tgClient,
		user:     user,
	}
}

func (m CurrentGameResult) Respond(ctx context.Context, message string) error {
	switch message {
	case nextInCurrentGameMessage:
		err := chooseNextRoundSuggestion(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseNextRoundSuggestion in CurrentGameResult: %w", err)
		}
		return nil
	case startAnewMessage:
		err := chooseNewStart0(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseNewStart0 in CurrentGameResult: %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in CurrentGameResult", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in CurrentGameResult: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in CurrentGameResult", message)
	}
}

func chooseCurrentGameResult(ctx context.Context, client *telegram.Client, user *userEntity.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.CurrentGameResult)
	if err != nil {
		return fmt.Errorf("failed in chooseEndGameResult changing current menu: %w", err)
	}
	thisMenu := NewCurrentGameResult(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseEndGameResult: %w", err)
	}
	return nil
}

func (m CurrentGameResult) sendDefaultMessage(ctx context.Context) error {
	gameResult := m.user.GameResult()
	result := "Текущие результаты игры:\n" + gameDetails(gameResult)
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		result,
		tgTypes.KeyboardButtonsFromStrings([]string{nextInCurrentGameMessage, startAnewMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func gameDetails(gameResult []userConstant.TeamInfo) string {
	var result string
	for _, teamInfo := range gameResult {
		result += fmt.Sprintf("\nКоманда %s:\n", teamInfo.Name)

		if len(teamInfo.RoundResults) == 0 {
			result += "\nЕще не было раундов\n"
		}

		for i, roundResult := range teamInfo.RoundResults {
			result += fmt.Sprintf("Раунд %d)  ✅%d   ❌%d   ❔%d\n", i+1, roundResult.CorrectAnswersCount, roundResult.IncorrectAnswersCount, roundResult.SkippedAnswersCount)
		}

		if len(teamInfo.RoundResults) > 1 {
			result += fmt.Sprintf(
				"\nИтог за все раунды: ✅%d   ❌%d   ❔%d   (раундов %d)\n",
				teamInfo.TotalCorrectAnswersCount,
				teamInfo.TotalIncorrectAnswersCount,
				teamInfo.TotalSkippedAnswersCount,
				len(teamInfo.RoundResults),
			)
		}
	}
	return result
}
