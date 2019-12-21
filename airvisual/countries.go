package airvisual

import (
	"context"
	"fmt"
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
	api_key := ctx.Value("API_KEY")
	if api_key == nil {
		return nil, nil, nil
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s%s", "v2/countries?key=", api_key))
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
