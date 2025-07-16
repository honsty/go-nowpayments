package payments

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/honsty/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// InvoiceArgs are the arguments used to make an invoice.
type InvoiceArgs struct {
	PaymentAmount

	CancelURL     string `json:"cancel_url,omitempty"`
	SuccessURL    string `json:"success_url,omitempty"`
	FeePaidByUser bool   `json:"is_fee_paid_by_user,omitempty"`
	FixedRate     bool   `json:"fixed_rate,omitempty"`
}

// Invoice describes an invoice. InvoiceURL is the URL to follow to make the payment.
// FIXME: inconsistency on their side for PriceAmount field should be a float64, like the field used for a payment not string
type Invoice struct {
	InvoiceArgs

	PriceAmount      string  `json:"price_amount"`
	ID               string  `json:"id"`
	CreatedAt        string  `json:"created_at,omitempty"`
	InvoiceURL       string  `json:"invoice_url,omitempty"`
	UpdatedAt        string  `json:"updated_at,omitempty"`
	TokenID          string  `json:"token_id,omitempty"`
	IsFixedRate      bool    `json:"is_fixed_rate,omitempty"`
	IsFeePaidByUser  bool    `json:"is_fee_paid_by_user,omitempty"`
	PayCurrency      *string `json:"pay_currency,omitempty"`
	IpnCallbackURL   string  `json:"ipn_callback_url,omitempty"`
	PartiallyPaidURL *string `json:"partially_paid_url,omitempty"`
	PayoutCurrency   *string `json:"payout_currency,omitempty"`
}

// NewInvoice creates an invoice
func NewInvoice(ia *InvoiceArgs) (*Invoice, error) {
	if ia == nil {
		return nil, errors.New("nil invoice args")
	}

	d, err := json.Marshal(ia)
	if err != nil {
		return nil, eris.Wrap(err, "invoice args")
	}

	p := &Invoice{}
	par := &core.SendParams{
		RouteName: "invoice-create",
		Into:      &p,
		Body:      strings.NewReader(string(d)),
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return p, nil
}
