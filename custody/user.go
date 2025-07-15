package custody

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

type ListCommonOptionsArgs struct {
	Id     int64
	Limit  int64
	Offset int64
	Order  string
}

type UserAccountArgs struct {
	Name string `json:"name"`
}

// User hold response for a Custody user account
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Balances hold multiple balances for Custody user account
type Balances struct {
	Usddtrc20 BalanceAmounts `json:"usddtrc20"`
	Usdtbsc   BalanceAmounts `json:"usdtbsc"`
}

// BalanceAmounts single balance for Custody user account
type BalanceAmounts struct {
	Amount        float64 `json:"amount"`
	PendingAmount float64 `json:"pendingAmount"`
}

// UserBalances hold response for a specific Custody user account balance
type UserBalances struct {
	SubPartnerID string   `json:"subPartnerId"`
	Balances     Balances `json:"balances"`
}

// NewUser will initiate new user account from a unique ID
// JWT is required for this request
func NewUser(cu *UserAccountArgs) (*User, error) {
	if cu == nil {
		return nil, errors.New("nil custody user account args")
	}

	d, err := json.Marshal(cu)
	if err != nil {
		return nil, eris.Wrap(err, "custody user account args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "custody user")
	}

	// CONSISTENCY PROBLEM ON THEIR SIDE: for some requests response is put under result object
	us := &core.V2ResponseFormat[*User]{}
	par := &core.SendParams{
		RouteName: "custody-create-account",
		Into:      &us,
		Body:      strings.NewReader(string(d)),
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return us.Result, nil
}

// ListUsers return a list of users based on filters provided in params
// JWT is required for this request

func ListUsers(o *ListCommonOptionsArgs) ([]*User, error) {
	u := url.Values{}

	if o != nil {
		if o.Id != 0 {
			u.Set("id", fmt.Sprintf("%d", o.Id))
		}
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		if o.Offset != 0 {
			u.Set("offset", fmt.Sprintf("%d", o.Offset))
		}
		if o.Order != "" {
			u.Set("order", o.Order)
		}
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list users")
	}

	usl := &core.V2ResponseFormat[[]*User]{}
	par := &core.SendParams{
		RouteName: "custody-list-users",
		Into:      usl,
		Values:    u,
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return usl.Result, nil
}

// GetBalance get the balances for a specific Custody user account, based on it's unique account ID
// This endpoint will work only if IP is whitelisted (or white IP restrictions are disabled)
func GetBalance(userAccountID string) (*UserBalances, error) {
	if userAccountID == "" {
		return nil, eris.New("empty user account ID")
	}

	bl := &core.V2ResponseFormat[*UserBalances]{}
	par := &core.SendParams{
		RouteName: "custody-account-balance",
		Path:      userAccountID,
		Into:      &bl,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return bl.Result, nil
}
