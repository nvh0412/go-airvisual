package airvisual

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.airvisual.com"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	common service

	Countries *CountryService
}

type service struct {
	client *Client
}

func NewClient() *Client {
	httpClient := &http.Client{}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.common.client = c
	c.Countries = (*CountryService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	defer resp.Body.Close()

	if v != nil {
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}

	return resp, err
}

type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"` // error message
}
