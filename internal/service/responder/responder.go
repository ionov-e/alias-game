package responder

import (
	userEntity "alias-game/internal/entity/user"
	"alias-game/internal/service/responder/helper"
	"alias-game/internal/storage"
	"alias-game/pkg/telegram"
	tgTypes "alias-game/pkg/telegram/types"
	"context"
	"fmt"
)

type Responder struct {
	tgUpdate     tgTypes.Update
	tgClient     telegram.Client
	userDB       storage.UserDBInterface
	dictionaryDB storage.DictionaryDBInterface
}

func New(tgUpdate tgTypes.Update, tgClient telegram.Client, userDB storage.UserDBInterface, dictionaryDB storage.DictionaryDBInterface) Responder {
	return Responder{
		tgUpdate:     tgUpdate,
		tgClient:     tgClient,
		userDB:       userDB,
		dictionaryDB: dictionaryDB,
	}
}

func (r *Responder) Run(ctx context.Context) error {
	tgUser, text, err := helper.ExtractFromUpdate(r.tgUpdate)
	if err != nil {
		return fmt.Errorf("failed at extracting from tgUpdate: %+v, error: %w", r.tgUpdate, err)
	}

	user, err := userEntity.NewFromTelegramUser(ctx, r.userDB, r.dictionaryDB, &tgUser)
	if err != nil {
		return fmt.Errorf("error getting user from Update.CallbackQuery: %w", err)
	}

	currentMenu, err := helper.MenuFactory(r.tgClient, &user)
	if err != nil {
		return fmt.Errorf("error getting choice from CallbackQuery.Message.Text: %w", err)
	}

	if err = currentMenu.Respond(ctx, text); err != nil {
		return fmt.Errorf("failed at responding to tgUpdate: %+v, error: %w", r.tgUpdate, err)
	}

	return nil
}
