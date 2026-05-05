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

func TestPhoneAPIKeyService_Store(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.PhoneAPIKeyStoreResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	phoneAPIKey, response, err := client.PhoneAPIKeys.Store(context.Background(), &PhoneAPIKeyStoreParams{
		Name: "test-key",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(phoneAPIKey)
	assert.JSONEq(t, string(stubs.PhoneAPIKeyStoreResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestPhoneAPIKeyService_StoreWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.HttpInternalServerErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.PhoneAPIKeys.Store(context.Background(), &PhoneAPIKeyStoreParams{
		Name: "test-key",
	})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.HttpInternalServerErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}

func TestPhoneAPIKeyService_StoreRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	var capturedRequest http.Request
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.PhoneAPIKeyStoreResponse(), &capturedRequest)
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, _, err := client.PhoneAPIKeys.Store(context.Background(), &PhoneAPIKeyStoreParams{
		Name: "test-key",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.MethodPost, capturedRequest.Method)
	assert.Equal(t, "/v1/phone-api-keys", capturedRequest.URL.Path)
	assert.Equal(t, "application/json", capturedRequest.Header.Get("Content-Type"))
	assert.Equal(t, apiKey, capturedRequest.Header.Get("x-api-key"))

	// Teardown
	server.Close()
}
