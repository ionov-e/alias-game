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
	token  string
	client http.Client
}

func New(token string) Client {
	return Client{
		token:  token,
		client: http.Client{},
	}
}

func (c *Client) SendMessage(chatId string, text string) (MessageResponse, error) {

	var responseObject MessageResponse
	data, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text":    text,
	})
	responseBytes, err := c.sendRequest("sendMessage", data)
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

func (c *Client) sendRequest(method string, data []byte) (responseBytes []byte, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   path.Join("bot"+c.token, method),
	}

	response, err := c.client.Post(u.String(), "application/json", bytes.NewReader(data))
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
