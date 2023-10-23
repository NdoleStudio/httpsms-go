package httpsms

import (
	"time"

	"github.com/google/uuid"
)

// MessageThreadIndexParams is the payload fetching entities.MessageThread sent between 2 numbers
type MessageThreadIndexParams struct {
	IsArchived bool    `json:"is_archived"`
	Skip       int     `json:"skip"`
	Query      *string `json:"query"`
	Limit      int     `json:"limit"`
	Owner      string  `json:"owner"`
}

// MessageThread represents a message thread between 2 phone numbers
type MessageThread struct {
	ID                 uuid.UUID `json:"id" example:"32343a19-da5e-4b1b-a767-3298a73703ca"`
	Owner              string    `json:"owner" example:"+18005550199"`
	Contact            string    `json:"contact" example:"+18005550100"`
	IsArchived         bool      `json:"is_archived" example:"false"`
	UserID             string    `json:"user_id" example:"WB7DRDWrJZRGbYrv2CKGkqbzvqdC"`
	Color              string    `json:"color" example:"indigo"`
	Status             string    `json:"status" example:"pending"`
	LastMessageContent string    `json:"last_message_content" example:"This is a sample message content"`
	LastMessageID      uuid.UUID `json:"last_message_id" example:"32343a19-da5e-4b1b-a767-3298a73703ca"`
	CreatedAt          time.Time `json:"created_at" example:"2022-06-05T14:26:09.527976+03:00"`
	UpdatedAt          time.Time `json:"updated_at" example:"2022-06-05T14:26:09.527976+03:00"`
	OrderTimestamp     time.Time `json:"order_timestamp"  example:"2022-06-05T14:26:09.527976+03:00"`
}

// MessageThreadsResponse is the response gotten with a message content
type MessageThreadsResponse ApiResponse[[]MessageThread]
