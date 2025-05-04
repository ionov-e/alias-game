package telegram

import (
	"alias-game/pkg/telegram/types"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time" //nolint:nolintlint,goimports
)

type Client struct {
	token  string
	log    *slog.Logger
	client http.Client
}

func NewClient(token string, log *slog.Logger) *Client {
	return &Client{
		token:  token,
		log:    log,
		client: http.Client{},
	}
}

func (c *Client) SendOneTimeReplyMarkup(ctx context.Context, chatID int64, text string, keyboardButtons [][]types.KeyboardButton) error {
	_, err := c.SendMessage(ctx, types.SendMessage{
		ChatID: chatID,
		Text:   text,
		ReplyMarkup: types.ReplyKeyboardMarkup{
			OneTimeKeyboard: true,
			Keyboard:        keyboardButtons,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send reply markup: chatID=%d, text=%s, keyboardButtons=%+v, error=%w", chatID, text, keyboardButtons, err)
	}
	return nil
}

func (c *Client) SendTextMessage(ctx context.Context, chatID int64, text string) (*types.MessageResponse, error) {
	messageResponse, err := c.SendMessage(ctx, types.SendMessage{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send text message: chatID=%d, text=%s, error=%w", chatID, text, err)
	}
	return messageResponse, nil
}

func (c *Client) SendMessage(ctx context.Context, message types.SendMessage) (*types.MessageResponse, error) {
	var messageResponse types.MessageResponse

	if message.ChatID == 0 {
		return nil, errors.New("chat id is required")
	}

	if message.Text == "" {
		return nil, errors.New("text is required")
	}

	data, err := message.Bytes()
	if err != nil {
		return nil, fmt.Errorf("SendMessage Bytes failed: %w", err)
	}

	responseBytes, err := c.sendRequest(ctx, "sendMessage", data)
	if err != nil {
		return nil, fmt.Errorf("send message failed: %w", err)
	}

	err = json.Unmarshal(responseBytes, &messageResponse)
	if err != nil {
		c.log.Info(fmt.Sprintf("Failed to unmarshal response: %s", string(responseBytes)))
		return &messageResponse, fmt.Errorf("failed to parse response: %w", err)
	}

	if !messageResponse.Ok {
		if messageResponse.Description != "" {
			return &messageResponse, errors.New("telegram error sending message: " + messageResponse.Description)
		}
		return &messageResponse, errors.New("telegram error sending message: No Description present in response")
	}

	return &messageResponse, nil
}

func (c *Client) UpdateMessageText(ctx context.Context, chatID int64, messageID int64, newText string) (bool, error) {
	messageToSend := types.EditMessageText{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      newText,
	}

	messageToSendInBytes, err := messageToSend.Bytes()
	if err != nil {
		return false, fmt.Errorf("UpdateMessageText Bytes failed: %w", err)
	}

	responseBytes, err := c.sendRequest(ctx, "editMessageText", messageToSendInBytes)
	if err != nil {
		return false, fmt.Errorf("edit message text request failed: %w", err)
	}

	// Parse the standard Telegram response
	var result struct {
		OK bool `json:"ok"`
		// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
		Result any `json:"result"`
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		return false, fmt.Errorf("failed to unmarshal editMessageText response: %w", err)
	}

	if !result.OK {
		return false, fmt.Errorf("telegram API returned ok=false: %s", string(responseBytes))
	}

	return true, nil
}

func (c *Client) DeleteMessage(ctx context.Context, chatID int64, messageID int64) (bool, error) {
	messageToSend := types.DeleteMessage{
		ChatID:    chatID,
		MessageID: messageID,
	}

	messageToSendInBytes, err := messageToSend.Bytes()
	if err != nil {
		return false, fmt.Errorf("DeleteMessage Bytes failed: %w", err)
	}

	responseBytes, err := c.sendRequest(ctx, "deleteMessage", messageToSendInBytes)
	if err != nil {
		return false, fmt.Errorf("delete message failed: %w", err)
	}

	var result struct {
		OK     bool `json:"ok"`
		Result bool `json:"result"`
	}

	if err := json.Unmarshal(responseBytes, &result); err != nil {
		return false, fmt.Errorf("failed to unmarshal deleteMessage response: %w", err)
	}

	if !result.OK {
		return false, fmt.Errorf("telegram API returned ok=false: %s", string(responseBytes))
	}

	return result.Result, nil
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
		c.log.Info(fmt.Sprintf("failed to parse response: %+v", err))
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return responseObject.Result, nil
}

func (c *Client) SetWebhook(ctx context.Context, webhookURL string) error {
	data, err := json.Marshal(map[string]string{"url": webhookURL})
	if err != nil {
		return fmt.Errorf("failed to marshal SetWebhook request: %w", err)
	}

	respBytes, err := c.sendRequest(ctx, "setWebhook", data)
	if err != nil {
		return fmt.Errorf("failed to send setWebhook request: %w", err)
	}

	var resp struct {
		OK          bool   `json:"ok"`
		Description string `json:"description,omitempty"`
	}
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to unmarshal setWebhook response: %w", err)
	}
	if !resp.OK {
		return fmt.Errorf("telegram API error: %s", resp.Description)
	}

	return nil
}

func (c *Client) DeleteWebhook(ctx context.Context) error {
	respBytes, err := c.sendRequest(ctx, "deleteWebhook", nil)
	if err != nil {
		return fmt.Errorf("failed to send setWebhook request: %w", err)
	}

	var resp struct {
		OK          bool   `json:"ok"`
		Description string `json:"description,omitempty"`
	}
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return fmt.Errorf("failed to unmarshal setWebhook response: %w", err)
	}
	if !resp.OK {
		return fmt.Errorf("telegram API error: %s", resp.Description)
	}

	return nil
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
		return nil, fmt.Errorf("failed to create sendRequest: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := c.client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("failed to send sendRequest: %w", err)
	}

	responseBytes, err = io.ReadAll(response.Body)
	defer func() { _ = response.Body.Close() }()
	if err != nil {
		return nil, fmt.Errorf("failed to read sendRequest response: %w", err)
	}

	return responseBytes, nil
}
