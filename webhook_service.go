package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
)

// WebhookService is the API client for the webhook endpoints
type WebhookService service

// WebhookStoreParams is the request payload for creating a webhook
type WebhookStoreParams struct {
	SigningKey   string   `json:"signing_key"`
	URL          string   `json:"url"`
	PhoneNumbers []string `json:"phone_numbers"`
	Events       []string `json:"events"`
}

// Store adds a new webhook
func (service *WebhookService) Store(ctx context.Context, params *WebhookStoreParams) (*WebhookResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1/webhooks", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	webhook := new(WebhookResponse)
	if err = json.Unmarshal(*response.Body, webhook); err != nil {
		return nil, response, err
	}

	return webhook, response, nil
}
