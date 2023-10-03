package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/godofprodev/tally/bot/models"
	"io"
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

func (c Client) GetRequest(url string) *models.Guild {
	resp, err := c.client.Get(url)
	if err != nil {
		// Handle the error.
		fmt.Println(err)
		return nil
	}

	// Close the response body when we're done with it.
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	guild := models.Guild{}

	err = json.Unmarshal(body, &guild)
	if err != nil {
		return nil
	}

	return &guild
}

func (c Client) PatchRequest(url string, userId string) {
	patch := []map[string]interface{}{
		{
			"userId": userId,
		},
	}

	// Marshal the JSON patch document
	patchJSON, err := json.Marshal(patch)
	if err != nil {
		// Handle error
		return
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewReader(patchJSON))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode)
		return
	}

	// The resource has been updated successfully
	fmt.Println("Resource updated successfully")
}
