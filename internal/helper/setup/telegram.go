package setup

import (
	"errors"
	"flag"
	"github.com/joho/godotenv" //nolint:nolintlint,goimports
	"log/slog"
	"os" //nolint:nolintlint,goimports
)

func TelegramBotToken(logger *slog.Logger) (string, error) {
	token := flag.String("token-bot-token", "", "telegram bot token")
	flag.Parse()
	if *token != "" {
		return *token, nil
	}
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found")
	}
	tokenFromEnv, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if ok {
		return tokenFromEnv, nil
	}
	return "", errors.New("getting token failed")
}
