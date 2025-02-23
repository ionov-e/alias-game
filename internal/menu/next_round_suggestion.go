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

const nextRoundMessage = "Начать раунд"

type NextRoundSuggestion struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewNextRoundSuggestion(tgClient *telegram.Client, u *user.User) NextRoundSuggestion {
	return NextRoundSuggestion{
		tgClient: tgClient,
		user:     u,
	}
}

func (m NextRoundSuggestion) Respond(ctx context.Context, message string) error {
	switch message {
	case nextRoundMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Word)
		if err != nil {
			return fmt.Errorf("failed in NextRoundSuggestion changing menu: %w", err)
		}
		newMenu := NewWordGuess(m.tgClient, m.user)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sendDefaultMessage in NextRoundSuggestion): %w", err)
		}
		return nil
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in NextRoundSuggestion", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in NextRoundSuggestion: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s', failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in NextRoundSuggestion", message)
	}
}

func (m NextRoundSuggestion) sendDefaultMessage(ctx context.Context) error {
	teamName := m.user.CurrentTeamName()
	var msg string
	if teamName != "" {
		msg = fmt.Sprintf("Следующая команда: %s", teamName)
	} else {
		msg = "Приготовься"
	}
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		msg,
		tgTypes.KeyboardButtonsFromStrings([]string{nextRoundMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
