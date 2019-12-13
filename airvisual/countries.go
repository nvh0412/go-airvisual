package airvisual

import (
	"context"
	"net/http"
)

type CountryService service

type Country struct {
	Name string `json:"country"`
}

type Countries struct {
	Status *string    `json:"status"`
	Data   *[]Country `json:"data"`
}

func (c Countries) String() string {
	return Stringify(c)
}

func (s *CountryService) ListCountries(ctx context.Context) (*Countries, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "v2/countries?key=")
	if err != nil {
		return nil, nil, err
	}

	countries := &Countries{}

	resp, err := s.client.Do(ctx, req, countries)
	if err != nil {
		return nil, resp, err
	}

	return countries, resp, nil
}
