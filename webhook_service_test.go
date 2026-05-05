package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/NdoleStudio/httpsms-go/internal/helpers"
	"github.com/NdoleStudio/httpsms-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestWebhookService_Store(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.WebhookStoreResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	webhook, response, err := client.Webhooks.Store(context.Background(), &WebhookStoreParams{
		SigningKey:   "whsec_test_signing_key",
		URL:          "https://example.com/webhook",
		PhoneNumbers: []string{"+18005550199"},
		Events:       []string{EventTypeMessagePhoneReceived, EventTypeMessagePhoneSent},
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(webhook)
	assert.JSONEq(t, string(stubs.WebhookStoreResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestWebhookService_StoreWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.HttpInternalServerErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.Webhooks.Store(context.Background(), &WebhookStoreParams{
		SigningKey:   "whsec_test_signing_key",
		URL:          "https://example.com/webhook",
		PhoneNumbers: []string{"+18005550199"},
		Events:       []string{EventTypeMessagePhoneReceived},
	})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.HttpInternalServerErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}

func TestWebhookService_StoreRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	var capturedRequest http.Request
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.WebhookStoreResponse(), &capturedRequest)
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, _, err := client.Webhooks.Store(context.Background(), &WebhookStoreParams{
		SigningKey:   "whsec_test_signing_key",
		URL:          "https://example.com/webhook",
		PhoneNumbers: []string{"+18005550199"},
		Events:       []string{EventTypeMessagePhoneReceived, EventTypeMessagePhoneSent},
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.MethodPost, capturedRequest.Method)
	assert.Equal(t, "/v1/webhooks", capturedRequest.URL.Path)
	assert.Equal(t, "application/json", capturedRequest.Header.Get("Content-Type"))
	assert.Equal(t, apiKey, capturedRequest.Header.Get("x-api-key"))

	// Teardown
	server.Close()
}
