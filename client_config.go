package httpsms

import "net/http"

type clientConfig struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://api.httpsms.com",
	}
}
