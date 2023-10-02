package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c Client) PostRequest(data any) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := c.client.Post("http://localhost:8080/v1/guilds", "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
