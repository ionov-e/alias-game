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

const startNewGameMessage = "Начать новую игру"

type EndGameResult struct {
	tgClient *telegram.Client
	user     *userEntity.User
}

func NewEndGameResult(tgClient *telegram.Client, user *userEntity.User) EndGameResult {
	return EndGameResult{
		tgClient: tgClient,
		user:     user,
	}
}

func (m EndGameResult) Respond(ctx context.Context, message string) error {
	switch message {
	case startNewGameMessage:
		err := chooseNewStart0(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseNewStart0: %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in Start0", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in Start0: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in Start0", message)
	}
}

func chooseEndGameResult(ctx context.Context, client *telegram.Client, user *userEntity.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.EndGameResult)
	if err != nil {
		return fmt.Errorf("failed in chooseEndGameResult changing current menu: %w", err)
	}
	thisMenu := NewEndGameResult(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseEndGameResult: %w", err)
	}
	user.ClearGame(ctx)
	return nil
}

func (m EndGameResult) sendDefaultMessage(ctx context.Context) error {
	gameResult := m.user.GameResult()
	winners := findWinners(gameResult)

	var result string
	if len(winners) == 1 {
		result = fmt.Sprintf("Победитель: 🏆 %s\n", winners[0].Name)
	} else {
		result = "Победу делят команды:\n"
		for _, winner := range winners {
			result += fmt.Sprintf("🏆 %s\n", winner.Name)
		}
	}

	result += "\n\nРезультаты:\n\n" + gameDetails(gameResult)
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		result,
		tgTypes.KeyboardButtonsFromStrings([]string{startNewGameMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func findWinners(teams []userConstant.TeamInfo) []userConstant.TeamInfo {
	var winners []userConstant.TeamInfo
	var maxCorrect uint16

	for _, team := range teams {
		if team.TotalCorrectAnswersCount > maxCorrect {
			winners = []userConstant.TeamInfo{team}
			maxCorrect = team.TotalCorrectAnswersCount
		} else if team.TotalCorrectAnswersCount == maxCorrect { // Tie
			winners = append(winners, team)
		}
	}

	return winners
}
