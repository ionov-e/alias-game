package main

import (
	"flag"
	"github.com/joho/godotenv"
	"go_telegram_start/telegram"
	"log"
	"os"
	"time"
)

func main() {
	mustSetUpLogging()
	telegramClient := telegram.New(mustToken())
	_, err := telegramClient.SendMessage("-4594910803", "Hello World!")
	if err != nil {
		log.Fatal(err)
	}
}

func mustSetUpLogging() {
	t := time.Now()
	logDir := "log/" + t.Format("2006/01")
	if err := os.MkdirAll(logDir, 0777); err != nil {
		log.Fatal(err)
	}
	logPath := logDir + "/" + t.Format("02") + ".log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
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