package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func main() {
	token := mustToken()
	chatId := "-4594910803"
	_, err := sendMessage(token, chatId, "Hello World!")
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

// Message https://core.telegram.org/bots/api#message
type Message struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id2"`
	} `json:"result"`
}

func sendMessage(token string, chatId string, text string) (Message, error) {

	var responseObject Message
	responseBytes, err := sendRequest("sendMessage", token, chatId, text)
	if err != nil {
		return responseObject, fmt.Errorf("send message failed: %w", err)
	}

	err = json.Unmarshal(responseBytes, &responseObject)
	if err != nil {
		log.Println(err)
		return responseObject, fmt.Errorf("failed to parse response: %w", err)
	}

	return responseObject, nil
}

func sendRequest(method string, token string, chatId string, text string) (responseBytes []byte, err error) {
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text":    text,
	})

	basePath := "bot" + token

	u := url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   path.Join(basePath, method),
	}

	client := http.Client{}
	resp, err := client.Post(u.String(), "application/json", bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	responseBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	return responseBytes, nil
}
