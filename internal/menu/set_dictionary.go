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

type SetDictionary0 struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

const easy1DictionaryNameMessage = "Легкий словарь"
const backMessage = "Назад"

func NewSetDictionary0(tgClient *telegram.Client, u *user.User, log *slog.Logger) SetDictionary0 {
	return SetDictionary0{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m SetDictionary0) Respond(ctx context.Context, message string) error {
	switch message {
	case easy1DictionaryNameMessage:
		err := m.user.ChooseDictionary(ctx, user.Easy1)
		if err != nil {
			return fmt.Errorf("failed ChooseDictionary in SetDictionary0: %w", err)
		}
		err = m.user.ChangeCurrentMenu(ctx, menuConstant.SetTeamCountPredefined)
		if err != nil {
			return fmt.Errorf("failed in SetDictionary0 changing current menu: %w", err)
		}
		newMenu := NewSetTeamCountPredefined(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in SetDictionary0: %w", err)
		}
		return nil
	case backMessage:
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.Start0)
		if err != nil {
			return fmt.Errorf("failed in SetDictionary0 changing current menu: %w", err)
		}
		newMenu := NewStart0(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in SetDictionary0: %w", err)
		}
		return nil
	default:
		m.log.Debug("unknown command in SetDictionary0", "message", message, "user_id", m.user.TelegramID())
		_, err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetDictionary0: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send menu message in SetDictionary0: %w", message, err)
		}
		return nil
	}
}

func (m SetDictionary0) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		"Выбери набор слов",
		tgTypes.KeyboardButtonsFromStrings([]string{easy1DictionaryNameMessage, backMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message in SetDictionary: %w", err)
	}

	return nil
}
