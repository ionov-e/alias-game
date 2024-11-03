package setup

import (
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func Logging() error {
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

func Token() (string, error) {
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
