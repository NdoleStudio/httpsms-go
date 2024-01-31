package httpsms

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"testing"

	"github.com/NdoleStudio/httpsms-go/internal/helpers"
	"github.com/NdoleStudio/httpsms-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestMessageThreadService_Index(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, stubs.MessageThreadIndexResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	indexParams := &MessageThreadIndexParams{
		IsArchived: false,
		Skip:       0,
		Query:      nil,
		Limit:      10,
		Owner:      fromNumber,
	}

	// Act
	threads, response, err := client.MessageThreads.Index(context.Background(), indexParams)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	jsonContent, _ := json.Marshal(threads)
	assert.JSONEq(t, string(stubs.MessageThreadIndexResponse()), string(jsonContent))

	// Teardown
	server.Close()
}

func TestMessageThreadService_IndexWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.HttpInternalServerErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	_, response, err := client.MessageThreads.Index(context.Background(), &MessageThreadIndexParams{})

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Equal(t, string(stubs.HttpInternalServerErrorResponse()), string(*response.Body))

	// Teardown
	server.Close()
}

func TestMessageThreadService_Delete(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusOK, nil)
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	response, err := client.MessageThreads.Delete(context.Background(), uuid.New())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestMessageThreadService_DeleteWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	apiKey := "test-api-key"
	server := helpers.MakeTestServer(http.StatusInternalServerError, stubs.HttpInternalServerErrorResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))

	// Act
	response, err := client.MessageThreads.Delete(context.Background(), uuid.New())

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
