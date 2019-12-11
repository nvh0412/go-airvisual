package airvisual

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "api.airvisual.com"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	ContryService *CountryService
}

type service struct {
	client *Client
}

type Response struct {
	*http.Response
}

type ErrorResponse struct {
}
