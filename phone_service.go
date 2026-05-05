package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
)

// PhoneService is the API client for the phone endpoints
type PhoneService service

// PhoneUpsertParams is the request payload for creating/updating a phone
type PhoneUpsertParams struct {
	PhoneNumber              string  `json:"phone_number"`
	FcmToken                 *string `json:"fcm_token,omitempty"`
	MessagesPerMinute        *uint   `json:"messages_per_minute,omitempty"`
	MaxSendAttempts          *uint   `json:"max_send_attempts,omitempty"`
	MessageExpirationSeconds *uint   `json:"message_expiration_seconds,omitempty"`
	SIM                      *string `json:"sim,omitempty"`
}

// PhoneFCMTokenParams is the request for binding FCM token to a phone via phone API key
type PhoneFCMTokenParams struct {
	PhoneNumber string `json:"phone_number"`
	FcmToken    string `json:"fcm_token"`
	SIM         string `json:"sim,omitempty"`
}

// Upsert creates or updates a phone
func (service *PhoneService) Upsert(ctx context.Context, params *PhoneUpsertParams) (*PhoneResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPut, "/v1/phones", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	phone := new(PhoneResponse)
	if err = json.Unmarshal(*response.Body, phone); err != nil {
		return nil, response, err
	}

	return phone, response, nil
}

// UpsertFCMToken binds an FCM token to a phone via the phone API key
func (service *PhoneService) UpsertFCMToken(ctx context.Context, params *PhoneFCMTokenParams) (*PhoneResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPut, "/v1/phones/fcm-token", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	phone := new(PhoneResponse)
	if err = json.Unmarshal(*response.Body, phone); err != nil {
		return nil, response, err
	}

	return phone, response, nil
}
