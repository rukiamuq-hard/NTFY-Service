package discord

import (
	"Service/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type discordWebhook struct {
	Content string `json:"content"`
}

type DSService struct {
}

func New() *DSService {
	return &DSService{}
}

func (ds *DSService) Send(ctx context.Context, req models.NotificationLog) error {
	message := discordWebhook{
		Content: req.Message,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.NewRequestWithContext(ctx, http.MethodPost, req.Recipient, bytes.NewReader(data))
	if err != nil {
		return err
	}
	resp.Header.Set("Content-type", "application/json")

	res, err := http.DefaultClient.Do(resp)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("Server returned status:%d", res.StatusCode)
	}

	return nil
}
