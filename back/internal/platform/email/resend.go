package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Sender interface {
	SendVerification(context.Context, string, string) error
}

type Resend struct {
	apiKey, from, verifyURL string
	client                  *http.Client
}

func NewResend(apiKey, from, verifyURL string) *Resend {
	return &Resend{apiKey: apiKey, from: from, verifyURL: verifyURL, client: http.DefaultClient}
}

func (r *Resend) SendVerification(ctx context.Context, to, token string) error {
	if r.apiKey == "" || r.from == "" {
		return fmt.Errorf("resend is not configured")
	}
	body, err := json.Marshal(map[string]string{"from": r.from, "to": to, "subject": "Verify your HireEasy email", "html": fmt.Sprintf("<p>Verify your email: <a href=\"%s?token=%s\">Verify email</a></p>", r.verifyURL, token)})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.resend.com/emails", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+r.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("resend returned status %d", resp.StatusCode)
	}
	return nil
}
