package custody

import (
	"fmt"
	"net/url"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/honsty/go-nowpayments/payments"
	"github.com/rotisserie/eris"
)

// BalanceAmounts single balance for Custody user account
type ListPaymentsOption struct {
	Limit        int64
	Page         int64
	Id           int64
	PayCurrency  string
	Status       string
	SubPartnerID string
	DateFrom     string
	DateTo       string
	OrderBy      string
	SortBy       string
}

// ListPayments return all Custody Payments, based on provided filters (which can be nil)
// JWT is required for this request
func ListPayments(o *ListPaymentsOption) ([]*payments.Payment[string], error) {
	u := url.Values{}

	if o != nil {
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		u.Set("page", fmt.Sprintf("%d", o.Page))
		if o.Id != 0 {
			u.Set("id", fmt.Sprintf("%d", o.Id))
		}
		if o.PayCurrency != "" {
			u.Set("pay_currency", o.PayCurrency)
		}
		if o.Status != "" {
			u.Set("status", o.Status)
		}
		if o.SubPartnerID != "" {
			u.Set("sub_partner_id", o.SubPartnerID)
		}
		if o.DateFrom != "" {
			u.Set("date_from", o.DateFrom)
		}
		if o.DateTo != "" {
			u.Set("date_to", o.DateTo)
		}
		if o.SortBy != "" {
			u.Set("sort_by", o.SortBy)
		}
		if o.OrderBy != "" {
			u.Set("order_by", o.OrderBy)
		}
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list payments")
	}

	pal := &core.V2ResponseFormat[[]*payments.Payment[string]]{}
	par := &core.SendParams{
		RouteName: "custody-payment-list",
		Into:      pal,
		Values:    u,
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return pal.Result, nil
}
