package exrates

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL = "https://api.exchangerate.host/"
	dateFmt = "2006-01-02"
)

type dayResponse struct {
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Values map[string]float64 `json:"rates"`
	Error  string             `json:"error"`
}

type intervalResponse struct {
	Base  string                        `json:"base"`
	Start string                        `json:"start_at"`
	End   string                        `json:"end_at"`
	Days  map[string]map[string]float64 `json:"rates"`
	Error string                        `json:"error"`
}

func dayRates(url string, params url.Values) (*Rates, error) {
	// Send API request.
	response := &dayResponse{}
	if err := doRequest(url, params, response); err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	// Parse API response.
	date, err := time.Parse(dateFmt, response.Date)
	if err != nil {
		return nil, err
	}

	return &Rates{
		Base:   response.Base,
		Date:   date,
		Values: response.Values,
	}, nil
}

func makeParams(base string, currencies []string) url.Values {
	params := url.Values{}
	if base := strings.ToUpper(strings.TrimSpace(base)); base != "" {
		params.Add("base", base)
	}

	symbols := []string{}
	for _, currency := range currencies {
		symbol := strings.ToUpper(strings.TrimSpace(currency))
		if symbol != "" {
			symbols = append(symbols, symbol)
		}
	}

	if len(symbols) > 0 {
		params.Add("symbols", strings.Join(symbols, ","))
	}

	return params
}

func doRequest(url string, params url.Values, w interface{}) error {
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, params.Encode())
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(w)
}
