package payments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/CIDgravity/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// Estimate holds the estimated amount from one currency to another
type Estimate struct {
	CurrencyFrom    string  `json:"currency_from"`
	CurrencyTo      string  `json:"currency_to"`
	AmountFrom      float64 `json:"amount_from"`
	EstimatedAmount string  `json:"estimated_amount"`
}

// EstimatedPrice calculates the approximate price from one currency to another (can be fiat or cryptocurrency)
func EstimatedPrice(amount float64, currencyFrom, currencyTo string) (*Estimate, error) {
	if amount == 0 {
		return nil, eris.New("use a price greater than zero")
	}

	u := url.Values{}
	u.Set("amount", fmt.Sprintf("%f", amount))
	u.Set("currency_from", currencyFrom)
	u.Set("currency_to", currencyTo)
	e := &Estimate{}

	par := &core.SendParams{
		RouteName: "estimate",
		Into:      &e,
		Values:    u,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// LatestEstimate holds info about the last price estimation
type LatestEstimate struct {
	PaymentID      json.Number `json:"id"`
	TokenID        string      `json:"token_id"`
	PayAmount      float64     `json:"pay_amount"`
	ExpirationDate string      `json:"expiration_estimate_date"`
}

// RefreshEstimatedPrice gets the current estimate on the payment and update the current estimate
func RefreshEstimatedPrice(paymentID string) (*LatestEstimate, error) {
	if paymentID == "" {
		return nil, errors.New("missing paymentID")
	}

	e := &LatestEstimate{}

	par := &core.SendParams{
		RouteName: "last-estimate",
		Into:      &e,
		Path:      paymentID + "/update-merchant-estimate",
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return e, nil
}
