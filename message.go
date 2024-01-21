package httpsms

import (
	"time"

	"github.com/google/uuid"
)

// MessageSendParams is the request payload for sending a message
type MessageSendParams struct {
	Content   string `json:"content"`
	From      string `json:"from"`
	Encrypted bool   `json:"encrypted"`
	RequestID string `json:"request_id,omitempty"`
	To        string `json:"to"`
}

// MessageIndexParams is the payload fetching entities.Message sent between 2 numbers
type MessageIndexParams struct {
	Skip    int     `json:"skip"`
	Contact string  `json:"contact"`
	Owner   string  `json:"owner"`
	Query   *string `json:"query"`
	Limit   int     `json:"limit"`
}

// MessageResponse is the response gotten with a message content
type MessageResponse ApiResponse[Message]

// MessagesResponse is the response with multiple messages
type MessagesResponse ApiResponse[[]Message]

// Message represents and incoming or outgoing SMS message
type Message struct {
	ID        uuid.UUID `json:"id" example:"32343a19-da5e-4b1b-a767-3298a73703cb"`
	RequestID *string   `json:"request_id" example:"153554b5-ae44-44a0-8f4f-7bbac5657ad4"`
	Owner     string    `json:"owner" example:"+18005550199"`
	UserID    string    `json:"user_id" example:"WB7DRDWrJZRGbYrv2CKGkqbzvqdC"`
	Contact   string    `json:"contact" example:"+18005550100"`
	Content   string    `json:"content" example:"This is a sample text message"`
	Type      string    `json:"type" example:"mobile-terminated"`
	Status    string    `json:"status" example:"pending"`
	Encrypted bool      `json:"encrypted" example:"false"`

	// SIM is the SIM card to use to send the message
	// * SMS1: use the SIM card in slot 1
	// * SMS2: use the SIM card in slot 2
	SIM string `json:"sim" example:"SIM1"`

	// SendDuration is the number of nanoseconds from when the request was received until when the mobile phone send the message
	SendDuration *int64 `json:"send_time" example:"133414"`

	RequestReceivedAt       time.Time  `json:"request_received_at" example:"2022-06-05T14:26:01.520828+03:00"`
	CreatedAt               time.Time  `json:"created_at" example:"2022-06-05T14:26:02.302718+03:00"`
	UpdatedAt               time.Time  `json:"updated_at" example:"2022-06-05T14:26:10.303278+03:00"`
	OrderTimestamp          time.Time  `json:"order_timestamp" gorm:"index:idx_messages_order_timestamp" example:"2022-06-05T14:26:09.527976+03:00"`
	LastAttemptedAt         *time.Time `json:"last_attempted_at" example:"2022-06-05T14:26:09.527976+03:00"`
	NotificationScheduledAt *time.Time `json:"scheduled_at" example:"2022-06-05T14:26:09.527976+03:00"`
	SentAt                  *time.Time `json:"sent_at" example:"2022-06-05T14:26:09.527976+03:00"`
	DeliveredAt             *time.Time `json:"delivered_at" example:"2022-06-05T14:26:09.527976+03:00"`
	ExpiredAt               *time.Time `json:"expired_at" example:"2022-06-05T14:26:09.527976+03:00"`
	FailedAt                *time.Time `json:"failed_at" example:"2022-06-05T14:26:09.527976+03:00"`
	CanBePolled             bool       `json:"can_be_polled" example:"false"`
	SendAttemptCount        uint       `json:"send_attempt_count" example:"0"`
	MaxSendAttempts         uint       `json:"max_send_attempts" example:"1"`
	ReceivedAt              *time.Time `json:"received_at" example:"2022-06-05T14:26:09.527976+03:00"`
	FailureReason           *string    `json:"failure_reason" example:"UNKNOWN"`
}
