package main

import (
	"flag"
	"github.com/joho/godotenv"
	"go_telegram_start/telegram"
	"log"
	"os"
)

func main() {
	telegramClient := telegram.New(mustToken())
	_, err := telegramClient.SendMessage("-4594910803", "Hello World!")
	if err != nil {
		log.Fatal(err)
	}
}

// Ensures that token is provided via flag or .env
func mustToken() string {
	token := flag.String("token-bot-token", "", "telegram bot token")
	flag.Parse()
	if *token == "" {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
		tokenFromEnv, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")
		if ok {
			return tokenFromEnv
		}

		log.Fatal("token is required")
	}
	return *token
}
