exrates
=======

[![Build Status](https://github.com/adrg/exrates/workflows/CI/badge.svg)](https://github.com/adrg/exrates/actions?query=workflow%3ACI)
[![pkg.go.dev documentation](https://pkg.go.dev/badge/github.com/adrg/exrates)](https://pkg.go.dev/github.com/adrg/exrates)
[![MIT license](http://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](http://opensource.org/licenses/MIT)
[![Go report card](https://goreportcard.com/badge/github.com/adrg/exrates)](https://goreportcard.com/report/github.com/adrg/exrates)

Exchange rates client for the excellent [Exchange rates API](https://exchangerate.host).
The API provides current and historical foreign exchange rates published by
financial data providers and banks, including the [European Central Bank](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html).

Full documentation can be found at: https://pkg.go.dev/github.com/adrg/exrates.

## Installation
```
go get github.com/adrg/exrates
```

## Usage

### Latest exchange rates

```go
// Get all available exchange rates.
rates, err := exrates.Latest("USD", nil)
if err != nil {
    // Treat error.
}
// Get specific exchange rates.
// rates, err := exrates.Latest("EUR", []string{"USD", "CAD"})

fmt.Printf("Exchange rates for %s on %s\n", rates.Base, rates.Date)
for currency, value := range rates.Values {
    fmt.Printf("%s: %f\n", currency, value)
}
```

### Exchange rates on specific date

```go
date := time.Date(2019, 3, 8, 0, 0, 0, 0, time.UTC)

// Get all available exchange rates.
rates, err := exrates.On("USD", date, nil)
if err != nil {
    // Treat error.
}
// Get specific exchange rates.
// rates, err := exrates.On("EUR", date, []string{"USD", "CAD"})

fmt.Printf("Exchange rates for %s on %s\n", rates.Base, rates.Date)
for currency, value := range rates.Values {
    fmt.Printf("%s: %f\n", currency, value)
}
```

### Exchange rates in date interval

```go
start := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2019, 4, 22, 0, 0, 0, 0, time.UTC)

// Get all available exchange rates.
days, err := exrates.Between("USD", start, end, nil)
if err != nil {
    // Treat error.
}
// Get specific exchange rates.
// days, err := exrates.Between("EUR", start, end, []string{"USD", "CAD"})

for _, day := range days {
    fmt.Printf("Exchange rates for %s on %s\n", day.Base, day.Date)
    for currency, value := range day.Values {
        fmt.Printf("%s: %f\n", currency, value)
    }
}
```

## Supported currencies

A list of supported currencies can be found [here](https://api.exchangerate.host/symbols?format=csv).

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.  
Before making a contribution please read [CONTRIBUTING.md](https://github.com/adrg/exrates/blob/master/CONTRIBUTING.md).

## References
For more information see the [Exchange rates API](https://exchangerate.host).

## License
Copyright (c) 2019 Adrian-George Bostan.  
This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
See [LICENSE](https://github.com/adrg/exrates/blob/master/LICENSE) for more details.
