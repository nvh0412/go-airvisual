package airvisual

import (
	"context"
	"fmt"
	"net/http"
)

type StateService service

type State struct {
	Name string `json:"state"`
}

type States struct {
	Status *string  `json:"status"`
	Data   *[]State `json:"data"`
}

func (s States) String() string {
	return Stringify(s)
}

func (s *StateService) ListStates(ctx context.Context, country string) (*States, *http.Response, error) {
	api_key := ctx.Value("API_KEY")
	if api_key == nil {
		return nil, nil, nil
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s%s%s%s", "v2/states?country=", country, "&key=", api_key))
	if err != nil {
		return nil, nil, err
	}

	states := &States{}

	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}
