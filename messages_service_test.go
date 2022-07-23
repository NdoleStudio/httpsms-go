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

const (
	fromNumber     = "+18005550199"
	toNumber       = "+18005550100"
	messageContent = "This is a sample text message"
)

func TestMessagesService_Send(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.MessagesSendResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	sendParams := &MessageSendParams{
		Content: messageContent,
		From:    fromNumber,
		To:      toNumber,
	}

	// Act
	message, response, err := client.Messages.Send(context.Background(), sendParams)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(message)
	assert.JSONEq(t, string(stubs.MessagesSendResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestMessagesService_SendWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.MessagesSendErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.Messages.Send(context.Background(), &MessageSendParams{})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.MessagesSendErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}
