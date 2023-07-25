package httpsms

import (
	"time"

	"github.com/google/uuid"
)

// Heartbeat represents is a pulse from an active phone
type Heartbeat struct {
	ID        uuid.UUID `json:"id" example:"32343a19-da5e-4b1b-a767-3298a73703cb"`
	Owner     string    `json:"owner" gorm:"index:idx_heartbeats_owner_timestamp" example:"+18005550199"`
	UserID    string    `json:"user_id" example:"WB7DRDWrJZRGbYrv2CKGkqbzvqdC"`
	Timestamp time.Time `json:"timestamp" example:"2022-06-05T14:26:01.520828+03:00"`
}

// HeartbeatIndexParams is the payload for fetching entities.Heartbeat of a phone number
type HeartbeatIndexParams struct {
	Skip  int     `json:"skip"`
	Owner string  `json:"owner"`
	Query *string `json:"query"`
	Limit int     `json:"limit"`
}

// HeartbeatsResponse is the response gotten with a message content
type HeartbeatsResponse ApiResponse[[]Heartbeat]
