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

const defaultSetRoundTimeMessage = "Выбор времени раунда"
const oneMinuteChoiceMessage = "1 минута"
const twoMinutesChoiceMessage = "2 минуты"
const threeMinutesChoiceMessage = "3 минуты"

type SetRoundTimePredefined struct {
	tgClient *telegram.Client
	user     *userEntity.User
}

func NewSetRoundTimePredefined(tgClient *telegram.Client, user *userEntity.User) SetRoundTimePredefined {
	return SetRoundTimePredefined{
		tgClient: tgClient,
		user:     user,
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

func (m SetRoundTimePredefined) setRoundTimeAndGoToNextMenu(ctx context.Context, newRoundTimeInSeconds uint16, message string) error {
	err := m.user.SetRoundTime(ctx, newRoundTimeInSeconds)
	if err != nil {
		return fmt.Errorf("error chooseDictionaryChoice0 (for message %s): %w", message, err)
	}
	err = chooseDictionaryChoice0(ctx, m.tgClient, m.user)
	if err != nil {
		return fmt.Errorf("error chooseDictionaryChoice0: %w", err)
	}
	return nil
}

func chooseSetRoundTime(ctx context.Context, client *telegram.Client, user *userEntity.User) error {
	err := user.ChangeCurrentMenu(ctx, menuConstant.SetRoundTimePredefined)
	if err != nil {
		return fmt.Errorf("failed in chooseSetRoundTime changing current menu: %w", err)
	}
	thisMenu := NewSetRoundTimePredefined(client, user)
	err = thisMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseSetRoundTime: %w", err)
	}
	return nil
}

func (m SetRoundTimePredefined) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		defaultSetRoundTimeMessage,
		tgTypes.KeyboardButtonsFromStrings(m.options()),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}

func (m SetRoundTimePredefined) options() []string {
	return []string{
		oneMinuteChoiceMessage,
		twoMinutesChoiceMessage,
		threeMinutesChoiceMessage,
	}
}
