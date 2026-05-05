package httpsms

import "time"

// Webhook represents a webhook registered in the httpSMS API
type Webhook struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	URL          string    `json:"url"`
	SigningKey   string    `json:"signing_key"`
	PhoneNumbers []string  `json:"phone_numbers"`
	Events       []string  `json:"events"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// WebhookResponse is the response gotten with a webhook
type WebhookResponse ApiResponse[Webhook]
