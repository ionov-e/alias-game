package menu

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
	"log/slog"
)

const defaultSetRoundTimeMessage = "Выбор времени раунда"
const oneMinuteChoiceMessage = "1 минута"
const twoMinutesChoiceMessage = "2 минуты"
const threeMinutesChoiceMessage = "3 минуты"

type SetRoundTimePredefined struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewSetRoundTimePredefined(tgClient *telegram.Client, u *user.User, log *slog.Logger) SetRoundTimePredefined {
	return SetRoundTimePredefined{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m SetRoundTimePredefined) Respond(ctx context.Context, message string) error {
	switch message {
	case oneMinuteChoiceMessage:
		return m.setRoundTimeAndGoToNextMenu(ctx, 1*60, message)
	case twoMinutesChoiceMessage:
		return m.setRoundTimeAndGoToNextMenu(ctx, 2*60, message)
	case threeMinutesChoiceMessage:
		return m.setRoundTimeAndGoToNextMenu(ctx, 3*60, message)
	default:
		m.log.Debug("unknown command in SetRoundTimePredefined", "message", message, "user_id", m.user.TelegramID())
		_, err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetRoundTimePredefined: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in SetRoundTimePredefined, failed to send message: %w", message, err)
		}
		return nil
	}
}

func (m SetRoundTimePredefined) setRoundTimeAndGoToNextMenu(ctx context.Context, newRoundTimeInSeconds uint16, message string) error {
	err := m.user.SetRoundTime(ctx, newRoundTimeInSeconds)
	if err != nil {
		return fmt.Errorf("error chooseDictionaryChoice0 (for message %s): %w", message, err)
	}
	err = m.user.ChangeCurrentMenu(ctx, menuConstant.SetDictionary)
	if err != nil {
		return fmt.Errorf("failed in chooseDictionaryChoice0 changing current menu: %w", err)
	}
	newMenu := NewSetDictionary0(m.tgClient, m.user, m.log)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseDictionaryChoice0: %w", err)
	}
	return nil
}

func (m SetRoundTimePredefined) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		defaultSetRoundTimeMessage,
		tgTypes.KeyboardButtonsFromStrings([]string{oneMinuteChoiceMessage, twoMinutesChoiceMessage, threeMinutesChoiceMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message in SetRoundTimePredefined: %w", err)
	}
	return nil
}
