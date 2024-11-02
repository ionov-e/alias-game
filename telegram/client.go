package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	token  string
	client http.Client
}

func New(token string) Client {
	return Client{
		token:  token,
		client: http.Client{},
	}
}

func (c *Client) SendMessage(ctx context.Context, chatID int64, text string) (MessageResponse, error) {
	var responseObject MessageResponse
	data, err := json.Marshal(map[string]string{
		"chat_id": strconv.FormatInt(chatID, 10),
		"text":    text,
	})
	if err != nil {
		return responseObject, fmt.Errorf("marshal SendMessage data failed: %w", err)
	}

	responseBytes, err := c.sendRequest(ctx, "sendMessage", data)
	if err != nil {
		return responseObject, fmt.Errorf("send message failed: %w", err)
	}

	err = json.Unmarshal(responseBytes, &responseObject)
	if err != nil {
		log.Printf("Failed to unmarshal response: %s", string(responseBytes))
		return responseObject, fmt.Errorf("failed to parse response: %w", err)
	}

	if !responseObject.IsOk() {
		return responseObject, errors.New("telegram error sending message: " + responseObject.Description)
	}

	return responseObject, nil
}

// GetUpdates https://core.telegram.org/bots/api#getupdates
func (c *Client) GetUpdates(ctx context.Context, offset uint64, limit int, timeout int) ([]Update, error) {
	var responseObject UpdateResponse
	data, err := json.Marshal(map[string]string{
		"offset":  strconv.FormatUint(offset, 10),
		"limit":   strconv.Itoa(limit),
		"timeout": strconv.Itoa(timeout),
	})

	if err != nil {
		return nil, fmt.Errorf("marshal GetUpdates data failed: %w", err)
	}

	responseBytes, err := c.sendRequest(ctx, "getUpdates", data)
	if err != nil {
		return nil, fmt.Errorf("get updates failed: %w", err)
	}

	err = json.Unmarshal(responseBytes, &responseObject)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return responseObject.Result, nil
}

func (c *Client) sendRequest(ctx context.Context, method string, data []byte) (responseBytes []byte, err error) {
	ctx, _ = context.WithTimeout(ctx, 40*time.Second) //nolint:nolintlint,govet

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.telegram.org/bot"+c.token+"/"+method,
		bytes.NewReader(data),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := c.client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	responseBytes, err = io.ReadAll(response.Body)
	defer func() { _ = response.Body.Close() }()
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return responseBytes, nil
}
