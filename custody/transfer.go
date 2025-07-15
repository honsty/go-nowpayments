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

// ListTransfersOption are options applying to the list of transfers
type ListTransfersOptionArgs struct {
	ListCommonOptionsArgs
	Status string
}

type TransferArgs struct {
	FromID   string  `json:"from_id"`
	ToID     string  `json:"to_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Transfer struct {
	Id        string    `json:"id,omitempty"`
	FromSubID string    `json:"from_sub_id,omitempty"`
	ToSubID   string    `json:"to_sub_id,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Amount    string    `json:"amount,omitempty"`
	Currency  string    `json:"currency,omitempty"`
}

// NewTransfer will initiate a transfer between two user account
// JWT is required for this request
func NewTransfer(ta *TransferArgs) (*Transfer, error) {
	if ta == nil {
		return nil, errors.New("nil transfer args")
	}

	d, err := json.Marshal(ta)
	if err != nil {
		return nil, eris.Wrap(err, "transfer args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list")
	}
	tr := &core.V2ResponseFormat[*Transfer]{}
	par := &core.SendParams{
		RouteName: "custody-transfer-create",
		Into:      &tr,
		Body:      strings.NewReader(string(d)),
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return tr.Result, nil
}

// GetTransfer will return single transfer information based on the supplied transfer ID
// JWT is required for this request
func GetTransfer(transferID string) (*Transfer, error) {
	if transferID == "" {
		return nil, eris.New("empty transfer ID")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list")
	}

	tr := &core.V2ResponseFormat[*Transfer]{}
	par := &core.SendParams{
		RouteName: "custody-transfer-single",
		Path:      transferID,
		Into:      &tr,
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return tr.Result, nil
}

// Transfer with return a list of all transfers based on supplied options (which can be nil)
// JWT is required for this request
func ListTransfers(o *ListTransfersOptionArgs) ([]*Transfer, error) {
	u := url.Values{}

	if o != nil {
		if o.Id != 0 {
			u.Set("id", fmt.Sprintf("%d", o.Id))
		}
		if o.Status != "" {
			u.Set("status", o.Status)
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
		return nil, eris.Wrap(err, "list")
	}

	trl := &core.V2ResponseFormat[[]*Transfer]{}
	par := &core.SendParams{
		RouteName: "custody-list-transfers",
		Into:      trl,
		Values:    u,
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return trl.Result, nil
}
