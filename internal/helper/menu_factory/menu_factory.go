package menu_factory

import (
	menuConstant "alias-game/internal/constant/menu"
	"alias-game/internal/menu"
	"alias-game/internal/user"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
	"log/slog"
)

type MenuInterface interface {
	Respond(ctx context.Context, message string) error
}

func MenuFactory(tgClient *telegram.Client, u *user.User, log *slog.Logger) (MenuInterface, error) {
	menuKeyString := u.CurrentMenuKey()
	menuKey := menuConstant.Key(menuKeyString)

	switch menuKey {
	case menuConstant.Start0:
		return menu.NewStart0(tgClient, u, log), nil
	case menuConstant.SetRoundTimePredefined:
		return menu.NewSetRoundTimePredefined(tgClient, u, log), nil
	case menuConstant.SetDictionary:
		return menu.NewSetDictionary0(tgClient, u, log), nil
	case menuConstant.SetTeamCountPredefined:
		return menu.NewSetTeamCountPredefined(tgClient, u, log), nil
	case menuConstant.SetTeamName:
		return menu.NewSetTeamName(tgClient, u, log), nil
	case menuConstant.SetWordCountToWinPredefined:
		return menu.NewSetWordCountToWinPredefined(tgClient, u, log), nil
	case menuConstant.NextRoundSuggestion:
		return menu.NewNextRoundSuggestion(tgClient, u, log), nil
	case menuConstant.Word:
		return menu.NewWordGuess(tgClient, u, log), nil
	case menuConstant.RoundResult:
		return menu.NewRoundResult(tgClient, u, log), nil
	case menuConstant.CurrentGameResult:
		return menu.NewCurrentGameResult(tgClient, u, log), nil
	case menuConstant.EndGameResult:
		return menu.NewEndGameResult(tgClient, u, log), nil
	}
	return nil, fmt.Errorf("menu factory called with: '%s' - no match", menuKeyString)
}
