package httpsms

import (
	"context"
	"encoding/json"
	"github.com/NdoleStudio/httpsms-go/internal/helpers"
	"github.com/NdoleStudio/httpsms-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHeartbeatService_Index(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.HeartbeatIndexResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	heartbeats, response, err := client.Heartbeats.Index(context.Background(), &HeartbeatIndexParams{
		Skip:  0,
		Owner: "+18005550199",
		Query: nil,
		Limit: 100,
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(heartbeats)
	assert.JSONEq(t, string(stubs.HeartbeatIndexResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestHeartbeatService_IndexWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.MessagesSendErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.Heartbeats.Index(context.Background(), &HeartbeatIndexParams{
		Skip:  0,
		Owner: "+18005550199",
		Query: nil,
		Limit: 100,
	})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.MessagesSendErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}
