package helper

import (
	menuConstant "alias-game/internal/constant/menu"
	userEntity "alias-game/internal/entity/user"
	"alias-game/internal/menu"
	"alias-game/pkg/telegram"
	"context"
	"fmt"
)

type MenuInterface interface {
	Respond(ctx context.Context, message string) error
}

func MenuFactory(tgClient *telegram.Client, user *userEntity.User) (MenuInterface, error) {
	menuKeyString := user.CurrentMenuKey()
	menuKey := menuConstant.Key(menuKeyString)

	switch menuKey {
	case menuConstant.Start0:
		return menu.NewStart0(tgClient, user), nil
	case menuConstant.SetRoundTimePredefined:
		return menu.NewSetRoundTimePredefined(tgClient, user), nil
	case menuConstant.SetDictionary:
		return menu.NewSetDictionary0(tgClient, user), nil
	case menuConstant.SetTeamCountPredefined:
		return menu.NewSetTeamCountPredefined(tgClient, user), nil
	case menuConstant.SetTeamName:
		return menu.NewSetTeamName(tgClient, user), nil
	case menuConstant.SetWordCountToWinPredefined:
		return menu.NewSetWordCountToWinPredefined(tgClient, user), nil
	case menuConstant.NextRoundSuggestion:
		return menu.NewNextRoundSuggestion(tgClient, user), nil
	case menuConstant.Word:
		return menu.NewWordGuess(tgClient, user), nil
	case menuConstant.RoundResult:
		return menu.NewRoundResult(tgClient, user), nil
	case menuConstant.CurrentGameResult:
		return menu.NewCurrentGameResult(tgClient, user), nil
	case menuConstant.EndGameResult:
		return menu.NewEndGameResult(tgClient, user), nil
	}
	return nil, fmt.Errorf("menu factory called with: '%s' - no match", menuKeyString)
}
