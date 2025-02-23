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

const defaultSetTeamCountMessage = "Выбор количества команд"
const oneTeamMessage = "Без команд (соло)"
const twoTeamsMessage = "2 команды"
const threeTeamsMessage = "3 команды"

type SetTeamCountPredefined struct {
	tgClient *telegram.Client
	user     *user.User
}

func NewSetTeamCountPredefined(tgClient *telegram.Client, u *user.User) SetTeamCountPredefined {
	return SetTeamCountPredefined{
		tgClient: tgClient,
		user:     u,
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
		newMenu := NewNextRoundSuggestion(m.tgClient, m.user)
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
		errMessage := fmt.Sprintf("Неизвестная комманда: '%s'", message)
		log.Printf("%s for user: %d in SetTeamCountPredefined", errMessage, m.user.TelegramID())
		err := m.tgClient.SendTextMessage(ctx, m.user.TelegramID(), errMessage)
		if err != nil {
			return fmt.Errorf("unexpected message '%s', failed to send text message in SetTeamCountPredefined: %w", message, err)
		}
		err = m.sendDefaultMessage(ctx)
		if err != nil {
			return fmt.Errorf("unexpected answer '%s' in SetTeamCountPredefined, failed to send message: %w", message, err)
		}
		return fmt.Errorf("unexpected answer '%s' in StaSetTeamCountPredefinedrt0", message)
	}
}

func (m SetTeamCountPredefined) setTeamCountAndChooseNextMenu(ctx context.Context, teamCount uint16) error {
	m.user.SetTeamCount(teamCount)
	err := m.user.ChangeCurrentMenu(ctx, menuConstant.SetTeamName)
	if err != nil {
		return fmt.Errorf("failed in chooseSetTeamNameChoice changing current menu: %w", err)
	}
	newMenu := NewSetTeamName(m.tgClient, m.user)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseSetTeamNameChoice: %w", err)
	}
	return nil
}

func chooseSetTeamCount(ctx context.Context, client *telegram.Client, u *user.User) error {
	err := u.ChangeCurrentMenu(ctx, menuConstant.SetTeamCountPredefined)
	if err != nil {
		return fmt.Errorf("failed in chooseSetTeamCount changing current menu: %w", err)
	}
	newMenu := NewSetTeamCountPredefined(client, u)
	err = newMenu.sendDefaultMessage(ctx)
	if err != nil {
		return fmt.Errorf("failed sending message in chooseSetTeamCount: %w", err)
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
