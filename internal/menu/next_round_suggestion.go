package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	userEntity "alias-game/internal/entity/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log"
)

const nextRoundMessage = "Начать раунд"

type NextRoundSuggestion struct {
	tgClient *telegram.Client
	user     *userEntity.User
}

func NewNextRoundSuggestion(tgClient *telegram.Client, user *userEntity.User) NextRoundSuggestion {
	return NextRoundSuggestion{
		tgClient: tgClient,
		user:     user,
	}
}

func (m NextRoundSuggestion) Respond(ctx context.Context, message string) error {
	switch message {
	case nextRoundMessage:
		err := chooseWordGuess(ctx, m.tgClient, m.user)
		if err != nil {
			return fmt.Errorf("error chooseWordGuess: %w", err)
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

func chooseNextRoundSuggestion(ctx context.Context, client *telegram.Client, user *userEntity.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.NextRoundSuggestion)
	if err != nil {
		return fmt.Errorf("failed in chooseNextRoundSuggestion changing current menu: %w", err)
	}
	thisMenu := NewNextRoundSuggestion(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseNextRoundSuggestion: %w", err)
	}
	return nil
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
		tgTypes.KeyboardButtonsFromStrings(m.options()),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func (m NextRoundSuggestion) options() []string {
	return []string{
		nextRoundMessage,
	}
}
