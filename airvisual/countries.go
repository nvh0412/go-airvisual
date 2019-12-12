package airvisual

import (
	"context"
	"net/http"
)

type CountryService service

type Country struct {
	Name string `json:"country"`
}

func (s *CountryService) ListCountries(ctx context.Context) (*[]Country, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "countries")
	if err != nil {
		return nil, nil, err
	}

	countries := &[]Country{}

	resp, err := s.client.Do(ctx, req, countries)
	if err != nil {
		return nil, resp, err
	}

	return countries, resp, nil
}
