package exrates

import (
	"errors"
	"sort"
	"time"
)

// Rates contains exchange rate values for a specific base currency on a
// particular date.
type Rates struct {
	Base   string
	Date   time.Time
	Values map[string]float64
}

// Latest returns the latest exchange rates for the selected base currency.
// Specific exchange rates can be requested by using the currencies parameter.
// If currencies is nil, all available exchange rates are returned.
func Latest(base string, currencies []string) (*Rates, error) {
	return dayRates(baseURL+"latest", makeParams(base, currencies))
}

// On returns the exchange rates available on the specified date, for the
// selected base currency. Specific exchange rates can be requested by using the
// currencies parameter. If currencies is nil, all available exchange rates are
// returned.
func On(base string, date time.Time, currencies []string) (*Rates, error) {
	// Check input date.
	if date.IsZero() {
		return nil, errors.New("invalid date parameter")
	}

	return dayRates(baseURL+date.Format(dateFmt), makeParams(base, currencies))
}

// Between returns exchange rates per day for the selected base currency, in the
// specified interval. Specific exchange rates can be requested by using the
// currencies parameter. If currencies is nil, all available exchange rates are
// returned.
func Between(base string, start, end time.Time, currencies []string) ([]*Rates, error) {
	// Check input interval.
	if start.IsZero() {
		return nil, errors.New("invalid start date parameter")
	}
	if end.IsZero() {
		return nil, errors.New("invalid end date parameter")
	}
	if start.After(end) {
		return nil, errors.New("start date must be before end date")
	}

	// Make request parameters.
	params := makeParams(base, currencies)
	params.Add("start_at", start.Format(dateFmt))
	params.Add("end_at", end.Format(dateFmt))

	// Send API request.
	response := &intervalResponse{}
	if err := doRequest(baseURL+"history", params, response); err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}

	lenDays := len(response.Days)
	if lenDays == 0 {
		return nil, errors.New("no rates found for the selected period")
	}

	// Parse API response.
	rates := make([]*Rates, lenDays)

	i := 0
	for day, values := range response.Days {
		date, err := time.Parse(dateFmt, day)
		if err != nil {
			return nil, err
		}

		rates[i] = &Rates{
			Base:   response.Base,
			Date:   date,
			Values: values,
		}

		i++
	}

	// Sort rates by date in ascending order.
	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Date.Before(rates[j].Date)
	})

	return rates, nil
}
