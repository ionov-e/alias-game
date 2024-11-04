package main

import (
	"alias-game/internal/app"
	"alias-game/internal/setup"
	"log"
)

func main() {
	if err := setup.Logging(); err != nil {
		log.Fatal(err)
	}

	botToken, err := setup.Token()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(botToken); err != nil {
		log.Fatal(err)
	}
}
