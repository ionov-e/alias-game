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

const defaultSetTeamCountMessage = "Выбор количества команд"
const oneTeamMessage = "Без команд (соло)"
const twoTeamsMessage = "2 команды"
const threeTeamsMessage = "3 команды"

type SetTeamCountPredefined struct {
	tgClient *telegram.Client
	user     *user.User
	log      *slog.Logger
}

func NewSetTeamCountPredefined(tgClient *telegram.Client, u *user.User, log *slog.Logger) SetTeamCountPredefined {
	return SetTeamCountPredefined{
		tgClient: tgClient,
		user:     u,
		log:      log,
	}
}

func (m SetTeamCountPredefined) Respond(ctx context.Context, message string) error {
	switch message {
	case oneTeamMessage:
		m.user.SetTeamCount(1)
		err := m.user.ChangeCurrentMenu(ctx, menuConstant.NextRoundSuggestion)
		if err != nil {
			return fmt.Errorf("failed in SetTeamCountPredefined changing current menu: %w", err)
		}
		newMenu := NewNextRoundSuggestion(m.tgClient, m.user, m.log)
		err = newMenu.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed sending message in SetTeamCountPredefined: %w", err)
		}
		return nil
	case twoTeamsMessage:
		return m.setTeamCountAndChooseNextMenu(ctx, 2)
	case threeTeamsMessage:
		return m.setTeamCountAndChooseNextMenu(ctx, 3)
	//TODO suggest Manual input
	default:
		m.log.Debug("unknown command in SetTeamCountPredefined", "message", message, "user_id", m.user.TelegramID())
		_, err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), fmt.Sprintf("Неизвестная комманда: '%s'", message))
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetTeamCountPredefined: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in SetTeamCountPredefined, failed to send message: %w", message, err)
		}
		return nil
	}
}

func (m SetTeamCountPredefined) setTeamCountAndChooseNextMenu(ctx context.Context, teamCount uint16) error {
	m.user.SetTeamCount(teamCount)
	err := m.user.ChangeCurrentMenu(ctx, menuConstant.SetTeamName)
	if err != nil {
		return fmt.Errorf("failed in setTeamCountAndChooseNextMenu changing current menu: %w", err)
	}
	newMenu := NewSetTeamName(m.tgClient, m.user, m.log)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in setTeamCountAndChooseNextMenu: %w", err)
	}
	return nil
}

func (m SetTeamCountPredefined) sendDefaultMessage(ctx context.Context) error {
	err := m.tgClient.SendOneTimeReplyMarkup(
		ctx,
		m.user.TelegramID(),
		defaultSetTeamCountMessage,
		tgTypes.KeyboardButtonsFromStrings([]string{oneTeamMessage, twoTeamsMessage, threeTeamsMessage}),
	)
	if err != nil {
		return fmt.Errorf("failed sending message: %w", err)
	}
	return nil
}
