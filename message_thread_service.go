package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// MessageThreadService is the API client for the `/message-threads` endpoint
type MessageThreadService service

// Index returns a list of contacts which a phone number has communicated with (threads). It will be sorted by timestamp in descending order.
//
// API Docs: https://api.httpsms.com/index.html#/Channel%20Threads/get_message_threads
func (service *MessageThreadService) Index(ctx context.Context, params *HeartbeatIndexParams) (*MessageThreadsResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v1/messageThreads", nil)
	if err != nil {
		return nil, nil, err
	}

	q := request.URL.Query()
	q.Add("skip", strconv.Itoa(params.Skip))
	q.Add("owner", params.Owner)
	q.Add("limit", strconv.Itoa(params.Limit))

	if params.Query != nil {
		q.Add("query", *params.Query)
	}

	request.URL.RawQuery = q.Encode()

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	messageThreads := new(MessageThreadsResponse)
	if err = json.Unmarshal(*response.Body, messageThreads); err != nil {
		return nil, response, err
	}

	return messageThreads, response, nil
}
