package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
)

// messagesService is the API client for the `/` endpoint
type messagesService service

// Send adds a new SMS message to be sent by the android phone
//
// API Docs: https://api.httpsms.com/index.html#/Messages/post_messages_send
func (service *messagesService) Send(ctx context.Context, params *MessageSendParams) (*MessageResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1/messages/send", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	message := new(MessageResponse)
	if err = json.Unmarshal(*response.Body, message); err != nil {
		return nil, response, err
	}

	return message, response, nil
}
