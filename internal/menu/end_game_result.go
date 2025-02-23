package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

const startNewGameMessage = "Начать новую игру"

type EndGameResult struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewEndGameResult(tgClient *telegram.Client, u *user.User) EndGameResult {
	return EndGameResult{
		tgClient: tgClient,
		user:     u,
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

func chooseEndGameResult(ctx context.Context, client *telegram.Client, u *user.User) error {
	err := u.ChangeCurrentMenu(ctx, menuConstant.EndGameResult)
	if err != nil {
		return fmt.Errorf("failed in chooseEndGameResult changing current menu: %w", err)
	}
	thisMenu := NewEndGameResult(client, u)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseEndGameResult: %w", err)
	}
	err = u.ClearGame(ctx)
	if err != nil {
		return fmt.Errorf("failed clearing game: %w", err)
	}
	return nil
}

func (m EndGameResult) sendDefaultMessage(ctx context.Context) error {
	result := m.user.EndGameResult()
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
