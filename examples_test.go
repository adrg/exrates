package exrates_test

import (
	"fmt"
	"time"

	"github.com/adrg/exrates"
)

func ExampleLatest() {
	// Get all available exchange rates.
	rates, err := exrates.Latest("USD", nil)
	if err != nil {
		// Treat error.
		return
	}
	// Get specific exchange rates.
	// rates, err := exrates.Latest("EUR", []string{"USD", "CAD"})

	fmt.Printf("Exchange rates for %s on %s\n", rates.Base, rates.Date)
	for currency, value := range rates.Values {
		fmt.Printf("%s: %f\n", currency, value)
	}
}

func ExampleOn() {
	date := time.Date(2019, 3, 8, 0, 0, 0, 0, time.UTC)

	// Get all available exchange rates.
	rates, err := exrates.On("USD", date, nil)
	if err != nil {
		// Treat error.
		return
	}
	// Get specific exchange rates.
	// rates, err := exrates.On("EUR", date, []string{"USD", "CAD"})

	fmt.Printf("Exchange rates for %s on %s\n", rates.Base, rates.Date)
	for currency, value := range rates.Values {
		fmt.Printf("%s: %f\n", currency, value)
	}
}

func ExampleBetween() {
	start := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2019, 4, 22, 0, 0, 0, 0, time.UTC)

	// Get all available exchange rates.
	days, err := exrates.Between("USD", start, end, nil)
	if err != nil {
		// Treat error.
		return
	}
	// Get specific exchange rates.
	// days, err := exrates.Between("EUR", start, end, []string{"USD", "CAD"})

	for _, day := range days {
		fmt.Printf("Exchange rates for %s on %s\n", day.Base, day.Date)
		for currency, value := range day.Values {
			fmt.Printf("%s: %f\n", currency, value)
		}
	}
}
