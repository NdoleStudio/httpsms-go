package httpsms

import "time"

// MessageSendParams is the request payload for sending a message
type MessageSendParams struct {
	Content string `json:"content"`
	From    string `json:"from"`
	To      string `json:"to"`
}

// MessageResponse is the response gotten with a message content
type MessageResponse struct {
	Data    Message `json:"data"`
	Message string  `json:"message"`
	Status  string  `json:"status"`
}

// Message represents and incoming or outgoing SMS message
type Message struct {
	Contact           string    `json:"contact"`
	Content           string    `json:"content"`
	CreatedAt         time.Time `json:"created_at"`
	FailureReason     string    `json:"failure_reason"`
	ID                string    `json:"id"`
	LastAttemptedAt   time.Time `json:"last_attempted_at"`
	OrderTimestamp    time.Time `json:"order_timestamp"`
	Owner             string    `json:"owner"`
	ReceivedAt        time.Time `json:"received_at"`
	RequestReceivedAt time.Time `json:"request_received_at"`
	SendTime          int       `json:"send_time"`
	SentAt            time.Time `json:"sent_at"`
	Status            string    `json:"status"`
	Type              string    `json:"type"`
	UpdatedAt         time.Time `json:"updated_at"`
	UserID            string    `json:"user_id"`
}
