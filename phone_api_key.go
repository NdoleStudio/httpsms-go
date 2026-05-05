package httpsms

import "time"

// PhoneAPIKey represents a phone API key
type PhoneAPIKey struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	UserID       string    `json:"user_id"`
	UserEmail    string    `json:"user_email"`
	PhoneNumbers []string  `json:"phone_numbers"`
	PhoneIDs     []string  `json:"phone_ids"`
	APIKey       string    `json:"api_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PhoneAPIKeyResponse is the response gotten with a phone api key
type PhoneAPIKeyResponse ApiResponse[PhoneAPIKey]
