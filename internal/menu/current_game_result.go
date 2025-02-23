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

const nextInCurrentGameMessage = "Продолжить"
const startAnewMessage = "Прекратить текущую игру"

type CurrentGameResult struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewCurrentGameResult(tgClient *telegram.Client, u *user.User) CurrentGameResult {
	return CurrentGameResult{
		tgClient: tgClient,
		user:     u,
	}
}

func (m CurrentGameResult) Respond(ctx context.Context, message string) error {
	switch message {
	case nextInCurrentGameMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.NextRoundSuggestion)
		if err != nil {
			return fmt.Errorf("failed in CurrentGameResult changing current menu: %w", err)
		}
		newMenu := NewNextRoundSuggestion(m.tgClient, m.user)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in CurrentGameResult: %w", err)
		}
		return nil
	case startAnewMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Start0)
		if err != nil {
			return fmt.Errorf("failed in CurrentGameResult changing current menu: %w", err)
		}
		newMenu := NewStart0(m.tgClient, m.user)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in CurrentGameResult: %w", err)
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

func (m CurrentGameResult) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		m.user.CurrentGameResul(),
		tgTypes.KeyboardButtonsFromStrings([]string{nextInCurrentGameMessage, startAnewMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
