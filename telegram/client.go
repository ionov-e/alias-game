package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(token string) Client {
	return Client{
		host:     "api.telegram.org",
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) SendMessage(chatId string, text string) (Message, error) {

	var responseObject Message
	responseBytes, err := c.sendRequest("sendMessage", chatId, text)
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

func (c *Client) sendRequest(method string, chatId string, text string) (responseBytes []byte, err error) {
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text":    text,
	})

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	response, err := c.client.Post(u.String(), "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	responseBytes, err = io.ReadAll(response.Body)
	defer func() { _ = response.Body.Close() }()
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
