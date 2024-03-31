package http

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct{}

func (f *Client) CreateClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	return client
}

func Get(url string) (*http.Response, error) {
	fn := &Client{}
	client := fn.CreateClient()

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	return resp, nil
}
