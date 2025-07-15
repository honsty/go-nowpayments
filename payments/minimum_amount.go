package payments

import (
	"net/url"

	"github.com/honsty/go-nowpayments/core"
)

// CurrencyAmount has info about minimum payment amount for a specific pair
type CurrencyAmount struct {
	CurrencyFrom   string  `json:"currency_from"`
	CurrencyTo     string  `json:"currency_to"`
	Amount         float64 `json:"min_amount"`
	FiatEquivalent float64 `json:"fiat_equivalent"`
}

// MinimumAmount returns the minimum payment amount for a specific pair
// fiatEquivalent is an optional param used to get equivalent amount in fiat currency (usd for example)
func MinimumAmount(currencyFrom, currencyTo, fiatEquivalent string) (*CurrencyAmount, error) {
	u := url.Values{}
	u.Set("currency_from", currencyFrom)
	u.Set("currency_to", currencyTo)

	if fiatEquivalent != "" {
		u.Set("fiat_equivalent", fiatEquivalent)
	}

	e := &CurrencyAmount{}

	par := &core.SendParams{
		RouteName: "min-amount",
		Into:      &e,
		Values:    u,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return e, nil
}
