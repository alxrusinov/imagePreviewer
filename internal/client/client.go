package client

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: &http.Client{}}
}

func (cl *Client) GetWithHeaders(url *url.URL, header http.Header) ([]byte, error) {
	req := &http.Request{URL: url, Method: http.MethodGet, Header: header}

	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
