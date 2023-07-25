package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// MessageService is the API client for the `/` endpoint
type MessageService service

// Send adds a new SMS message to be sent by the android phone
//
// API Docs: https://api.httpsms.com/index.html#/Messages/post_messages_send
func (service *MessageService) Send(ctx context.Context, params *MessageSendParams) (*MessageResponse, *Response, error) {
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

// Index returns a list of messages which are sent between 2 phone numbers. It will be sorted by timestamp in descending order.
//
// API Docs: https://api.httpsms.com/index.html#/Messages/get_messages
func (service *MessageService) Index(ctx context.Context, params *MessageIndexParams) (*MessagesResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v1/messages", nil)
	if err != nil {
		return nil, nil, err
	}

	q := request.URL.Query()
	q.Add("skip", strconv.Itoa(params.Skip))
	q.Add("owner", params.Owner)
	q.Add("contact", params.Contact)
	q.Add("limit", strconv.Itoa(params.Limit))

	if params.Query != nil {
		q.Add("query", *params.Query)
	}

	request.URL.RawQuery = q.Encode()

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	messages := new(MessagesResponse)
	if err = json.Unmarshal(*response.Body, messages); err != nil {
		return nil, response, err
	}

	return messages, response, nil
}
