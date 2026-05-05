package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
)

// PhoneAPIKeyService is the API client for the phone api key endpoints
type PhoneAPIKeyService service

// PhoneAPIKeyStoreParams is the request payload for creating a phone api key
type PhoneAPIKeyStoreParams struct {
	Name string `json:"name"`
}

// Store adds a new phone api key
func (service *PhoneAPIKeyService) Store(ctx context.Context, params *PhoneAPIKeyStoreParams) (*PhoneAPIKeyResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1/phone-api-keys", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	phoneAPIKey := new(PhoneAPIKeyResponse)
	if err = json.Unmarshal(*response.Body, phoneAPIKey); err != nil {
		return nil, response, err
	}

	return phoneAPIKey, response, nil
}
