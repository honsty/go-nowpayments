package recurring_payments

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// ReccuringPaymentArgs handle args to create a recurring payment for a specific custody user account
type RecurringPaymentArgs struct {
	SubscriptionPlanID int64 `json:"subscription_plan_id"`
	SubPartnerID       int64 `json:"sub_partner_id"`
}

// RecurringPayment handle status of a specific recurring payment
type RecurringPayment struct {
	ID                 string     `json:"id"`
	SubscriptionPlanID string     `json:"subscription_plan_id"`
	IsActive           bool       `json:"is_active"`
	Status             string     `json:"status"`
	ExpireDate         string     `json:"expire_date"`
	Subscriber         Subscriber `json:"subscriber"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

// DeleteReccurringPayment handle status when deleting recurring payment
type DeleteReccurringPayment struct {
	Status string `json:"status"`
}

// Subscriber handle a subscriber to a specific plan
type Subscriber struct {
	Email        string `json:"email,omitempty"`
	SubPartnerID string `json:"sub_partner_id,omitempty"`
}

// New will create new recurring payment from custody user account
// This require an existing user account (created using custody.Create method)
// JWT is required for this request
func New(ru *RecurringPaymentArgs) (*RecurringPayment, error) {
	if ru == nil {
		return nil, errors.New("nil recurring payment args")
	}

	d, err := json.Marshal(ru)
	if err != nil {
		return nil, eris.Wrap(err, "recurring payment args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "recurring payment")
	}

	// Inconsistency on their side: single sub partner ID is allowed, but response is an array
	// will return only the first element of array
	rcu := &core.V2ResponseFormat[[]*RecurringPayment]{}
	par := &core.SendParams{
		RouteName: "recurring-payment-create",
		Into:      &rcu,
		JWTToken:  tok,
		Body:      strings.NewReader(string(d)),
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return rcu.Result[0], nil
}

// Get return a single reccuring payment via it's ID
func Get(recurringPaymentID string) (*RecurringPayment, error) {
	if recurringPaymentID == "" {
		return nil, eris.New("empty recurring payment ID")
	}

	rp := &core.V2ResponseFormat[*RecurringPayment]{}
	par := &core.SendParams{
		RouteName: "recurring-payment-single",
		Path:      recurringPaymentID,
		Into:      &rp,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return rp.Result, nil
}

// Delete remove a recurring payment via it's ID
// JWT is required for this request
func Delete(recurringPaymentID string) (*string, error) {
	if recurringPaymentID == "" {
		return nil, eris.New("empty recurring payment ID")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "recurring payment")
	}

	de := &core.V2ResponseFormat[*string]{}
	par := &core.SendParams{
		RouteName: "recurring-payment-delete",
		Path:      recurringPaymentID,
		Into:      &de,
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return de.Result, nil
}
