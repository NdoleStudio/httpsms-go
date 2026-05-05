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

func TestPhoneService_Upsert(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.PhoneUpsertResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	fcmToken := "test-fcm-token"
	messagesPerMinute := uint(5)
	maxSendAttempts := uint(3)
	messageExpiration := uint(3600)
	sim := "SIM1"
	phone, response, err := client.Phones.Upsert(context.Background(), &PhoneUpsertParams{
		PhoneNumber:              "+18005550199",
		FcmToken:                 &fcmToken,
		MessagesPerMinute:        &messagesPerMinute,
		MaxSendAttempts:          &maxSendAttempts,
		MessageExpirationSeconds: &messageExpiration,
		SIM:                      &sim,
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(phone)
	assert.JSONEq(t, string(stubs.PhoneUpsertResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestPhoneService_UpsertWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.HttpInternalServerErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.Phones.Upsert(context.Background(), &PhoneUpsertParams{
		PhoneNumber: "+18005550199",
	})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.HttpInternalServerErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}

func TestPhoneService_UpsertRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	var capturedRequest http.Request
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.PhoneUpsertResponse(), &capturedRequest)
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	fcmToken := "test-fcm-token"
	messagesPerMinute := uint(5)
	_, _, err := client.Phones.Upsert(context.Background(), &PhoneUpsertParams{
		PhoneNumber:       "+18005550199",
		FcmToken:          &fcmToken,
		MessagesPerMinute: &messagesPerMinute,
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.MethodPut, capturedRequest.Method)
	assert.Equal(t, "/v1/phones", capturedRequest.URL.Path)
	assert.Equal(t, "application/json", capturedRequest.Header.Get("Content-Type"))
	assert.Equal(t, apiKey, capturedRequest.Header.Get("x-api-key"))

	// Teardown
	server.Close()
}

func TestPhoneService_UpsertFCMToken(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.PhoneUpsertFCMTokenResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	phone, response, err := client.Phones.UpsertFCMToken(context.Background(), &PhoneFCMTokenParams{
		PhoneNumber: "+18005550199",
		FcmToken:    "new-fcm-token",
		SIM:         "SIM1",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(phone)
	assert.JSONEq(t, string(stubs.PhoneUpsertFCMTokenResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestPhoneService_UpsertFCMTokenWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.HttpInternalServerErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.Phones.UpsertFCMToken(context.Background(), &PhoneFCMTokenParams{
		PhoneNumber: "+18005550199",
		FcmToken:    "new-fcm-token",
	})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.HttpInternalServerErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}

func TestPhoneService_UpsertFCMTokenRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	var capturedRequest http.Request
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.PhoneUpsertFCMTokenResponse(), &capturedRequest)
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, _, err := client.Phones.UpsertFCMToken(context.Background(), &PhoneFCMTokenParams{
		PhoneNumber: "+18005550199",
		FcmToken:    "new-fcm-token",
		SIM:         "SIM1",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.MethodPut, capturedRequest.Method)
	assert.Equal(t, "/v1/phones/fcm-token", capturedRequest.URL.Path)
	assert.Equal(t, "application/json", capturedRequest.Header.Get("Content-Type"))
	assert.Equal(t, apiKey, capturedRequest.Header.Get("x-api-key"))

	// Teardown
	server.Close()
}
