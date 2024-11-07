package telegram

import (
	"alias-game/pkg/telegram/types"
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

func (c *Client) SendMessage(ctx context.Context, message types.SendMessage) (types.MessageResponse, error) {
	var messageResponse types.MessageResponse

	if message.ChatID == 0 {
		return messageResponse, errors.New("chat id is required")
	}

	if message.Text == "" {
		return messageResponse, errors.New("text is required")
	}

	data, err := message.ToJSON()
	if err != nil {
		return messageResponse, fmt.Errorf("SendMessage ToJSON failed: %w", err)
	}

	responseBytes, err := c.sendRequest(ctx, "sendMessage", data)
	if err != nil {
		return messageResponse, fmt.Errorf("send message failed: %w", err)
	}

	err = json.Unmarshal(responseBytes, &messageResponse)
	if err != nil {
		log.Printf("Failed to unmarshal response: %s", string(responseBytes))
		return messageResponse, fmt.Errorf("failed to parse response: %w", err)
	}

	if !messageResponse.IsOk() {
		return messageResponse, errors.New("telegram error sending message: " + messageResponse.DescriptionText())
	}

	return messageResponse, nil
}

// GetUpdates https://core.telegram.org/bots/api#getupdates
func (c *Client) GetUpdates(ctx context.Context, offset uint64, limit int, timeout int) ([]types.Update, error) {
	var responseObject types.UpdateResponse
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
