package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DiscordWebhookAction posts a plain content message to a Discord
// webhook URL. Just "content" for now — no embeds or attachments.
type DiscordWebhookAction struct {
	WebhookURL string
	Message    string
}

func (a *DiscordWebhookAction) Execute() error {
	body, err := json.Marshal(map[string]string{"content": a.Message})
	if err != nil {
		return err
	}

	resp, err := http.Post(a.WebhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("discord action: posted -> %s\n", resp.Status)
	return nil
}
