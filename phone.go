package httpsms

import "time"

// Phone represents a phone registered in the httpSMS API
type Phone struct {
	ID                       string    `json:"id"`
	UserID                   string    `json:"user_id"`
	PhoneNumber              string    `json:"phone_number"`
	FcmToken                 *string   `json:"fcm_token"`
	MessagesPerMinute        uint      `json:"messages_per_minute"`
	MaxSendAttempts          uint      `json:"max_send_attempts"`
	MessageExpirationSeconds uint      `json:"message_expiration_seconds"`
	SIM                      string    `json:"sim"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// PhoneResponse is the response gotten with a phone
type PhoneResponse ApiResponse[Phone]
