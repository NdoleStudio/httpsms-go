package httpsms

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HeartbeatService is the API client for the `/heartbeats` endpoint
type HeartbeatService service

// Index returns a list of heartbeats from an android phone. It will be sorted by timestamp in descending order.
//
// API Docs: https://api.httpsms.com/index.html#/Heartbeats/get_heartbeats
func (service *HeartbeatService) Index(ctx context.Context, params *HeartbeatIndexParams) (*HeartbeatsResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v1/heartbeats", nil)
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

	heartbeats := new(HeartbeatsResponse)
	if err = json.Unmarshal(*response.Body, heartbeats); err != nil {
		return nil, response, err
	}

	return heartbeats, response, nil
}
