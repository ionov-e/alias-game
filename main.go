package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"go_telegram_start/alias"
	"go_telegram_start/database"
	"go_telegram_start/telegram"
	"log"
	"os"
	"time"
)

func main() {
	if err := setUpLogging(); err != nil {
		log.Fatal(err)
	}

	token, err := token()
	if err != nil {
		log.Fatal(err)
	}

	telegramClient := telegram.New(token)
	localDB := database.NewLocal()
	ctx := context.Background()
	lastUpdateID := uint64(0)

	for {
		updates, err := telegramClient.GetUpdates(ctx, lastUpdateID, 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			if update.UpdateID == lastUpdateID {
				continue
			}
			lastUpdateID = update.UpdateID

			game := alias.NewGame(update, telegramClient, localDB)
			if err := game.Respond(ctx); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func setUpLogging() error {
	t := time.Now()
	logDir := "log/" + t.Format("2006/01")
	if err := os.MkdirAll(logDir, 0777); err != nil {
		return fmt.Errorf("creating log dir failed: %w", err)
	}
	logPath := logDir + "/" + t.Format("02") + ".log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("opening log file failed: %w", err)
	}
	log.SetOutput(file)
	return nil
}

// Ensures that token is provided via flag or .env
func token() (string, error) {
	token := flag.String("token-bot-token", "", "telegram bot token")
	flag.Parse()
	if *token != "" {
		return *token, nil
	}
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	tokenFromEnv, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if ok {
		return tokenFromEnv, nil
	}
	return "", errors.New("getting token failed")
}
