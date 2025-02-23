package menu_factory

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
)

type MenuInterface interface {
	Respond(ctx context.Context, message string) error
}

func MenuFactory(tgClient *telegram.Client, u *user.User) (MenuInterface, error) {
	menuKeyString := u.CurrentMenuKey()
	menuKey := menuConstant.Key(menuKeyString)

	switch menuKey {
	case menuConstant.Start0:
		return menu.NewStart0(tgClient, u), nil
	case menuConstant.SetRoundTimePredefined:
		return menu.NewSetRoundTimePredefined(tgClient, u), nil
	case menuConstant.SetDictionary:
		return menu.NewSetDictionary0(tgClient, u), nil
	case menuConstant.SetTeamCountPredefined:
		return menu.NewSetTeamCountPredefined(tgClient, u), nil
	case menuConstant.SetTeamName:
		return menu.NewSetTeamName(tgClient, u), nil
	case menuConstant.SetWordCountToWinPredefined:
		return menu.NewSetWordCountToWinPredefined(tgClient, u), nil
	case menuConstant.NextRoundSuggestion:
		return menu.NewNextRoundSuggestion(tgClient, u), nil
	case menuConstant.Word:
		return menu.NewWordGuess(tgClient, u), nil
	case menuConstant.RoundResult:
		return menu.NewRoundResult(tgClient, u), nil
	case menuConstant.CurrentGameResult:
		return menu.NewCurrentGameResult(tgClient, u), nil
	case menuConstant.EndGameResult:
		return menu.NewEndGameResult(tgClient, u), nil
	}
	return nil, fmt.Errorf("menu factory called with: '%s' - no match", menuKeyString)
}
