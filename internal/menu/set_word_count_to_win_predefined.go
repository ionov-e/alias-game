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

const defaultSetWordCountToWinPredefinedMessage = "Выбор количества слов для победы"
const oneHundredChoiceMessage = "100"
const twoHundredChoiceMessage = "200"
const threeHundredChoiceMessage = "300"

type SetWordCountToWinPredefined struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewSetWordCountToWinPredefined(tgClient *telegram.Client, u *user.User) SetWordCountToWinPredefined {
	return SetWordCountToWinPredefined{
		tgClient: tgClient,
		user:     u,
	}
}

func (m SetWordCountToWinPredefined) Respond(ctx context.Context, message string) error {
	switch message {
	case oneHundredChoiceMessage:
		return m.setWordCountToWinAndGoToNextMenu(ctx, 100)
	case twoHundredChoiceMessage:
		return m.setWordCountToWinAndGoToNextMenu(ctx, 200)
	case threeHundredChoiceMessage:
		return m.setWordCountToWinAndGoToNextMenu(ctx, 300)
	//TODO suggest Manual input
	default:
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in SetWordCountToWinPredefined", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetWordCountToWinPredefined: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in SetWordCountToWinPredefined, failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in SetWordCountToWinPredefined", message)
	}
}

func (m SetWordCountToWinPredefined) setWordCountToWinAndGoToNextMenu(ctx context.Context, wordCountToWin uint16) error {
	m.user.SetWordCountToWin(wordCountToWin)
	err := m.user.ChangeCurrentMenu(ctx, menuConstant.NextRoundSuggestion)
	if err != nil {
		return fmt.Errorf("failed in SetWordCountToWinPredefined changing current menu: %w", err)
	}
	newMenu := NewNextRoundSuggestion(m.tgClient, m.user)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in SetWordCountToWinPredefined: %w", err)
	}
	return nil
}

func (m SetWordCountToWinPredefined) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		defaultSetWordCountToWinPredefinedMessage,
		tgTypes.KeyboardButtonsFromStrings([]string{oneHundredChoiceMessage, twoHundredChoiceMessage, threeHundredChoiceMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
