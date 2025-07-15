package payments

import (
	"fmt"
	"net/url"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// ListOption are options applying to the list of transactions
type ListOption struct {
	DateFrom string
	DateTo   string
	Limit    int
	OrderBy  string
	Page     int
	SortBy   string
}

// List returns a list of all transactions, depending on the supplied options (which can be nil)
// JWT is required for this request
func List(o *ListOption) ([]*Payment[int64], error) {
	u := url.Values{}

	if o != nil {
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		if o.DateFrom != "" {
			u.Set("dateFrom", o.DateFrom)
		}
		if o.DateTo != "" {
			u.Set("dateTo", o.DateTo)
		}
		u.Set("page", fmt.Sprintf("%d", o.Page))
		if o.SortBy != "" {
			u.Set("sortBy", o.SortBy)
		}
		if o.OrderBy != "" {
			u.Set("orderBy", o.OrderBy)
		}
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list")
	}

	type plist struct {
		Data []*Payment[int64] `json:"data"`
	}

	pl := &plist{Data: make([]*Payment[int64], 0)}
	par := &core.SendParams{
		RouteName: "payments-list",
		Into:      pl,
		Values:    u,
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return pl.Data, nil
}
