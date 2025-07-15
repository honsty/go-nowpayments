package custody

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// NewWriteOffToMaster will initiate a funds transfer from user balance to master account
// JWT is required for this request
func NewWriteOffToMaster(wo *DepositArgs) (*Transfer, error) {
	if wo == nil {
		return nil, errors.New("nil write off args")
	}

	d, err := json.Marshal(wo)
	if err != nil {
		return nil, eris.Wrap(err, "write off args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "custody write-off to master")
	}

	tr := &core.V2ResponseFormat[*Transfer]{}
	par := &core.SendParams{
		RouteName: "custody-write-off-to-master",
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
