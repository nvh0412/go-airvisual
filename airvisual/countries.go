package airvisual

type CountryService service

type Country struct {
	Name string `json:"country"`
}

func (s *CountryService) ListCountries() ([]Country, *Response, error) {
	return nil, nil, nil
}
