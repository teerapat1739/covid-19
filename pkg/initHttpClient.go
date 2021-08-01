package pkg

import (
	"net/http"
	"time"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

var (
	Client HTTPClient
)

type httpClient struct {
	client *http.Client
}

func NewHttpClient() HTTPClient {
	return &httpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *httpClient) Get(url string) (*http.Response, error) {

	return h.client.Get(url)
}
