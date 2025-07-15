package subscriptions

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	recurringPayment "github.com/honsty/go-nowpayments/recurring_payments"
	"github.com/rotisserie/eris"
)

// SubscriptionArgs handle args to create a subscription plan
type SubscriptionArgs struct {
	Title       string  `json:"title,omitempty"`
	IntervalDay int64   `json:"interval_day,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Currency    string  `json:"currency,omitempty"`
}

// EmailSubscriptionArgs handle args to create a subscription with an email
type EmailSubscriptionArgs struct {
	SubscriptionPlanID int64  `json:"subscription_plan_id"`
	Email              string `json:"email"`
}

// Subscription handle subscription plan
type Subscription struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	IntervalDay      string    `json:"interval_day"`
	IpnCallbackURL   string    `json:"ipn_callback_url,omitempty"`
	SuccessURL       string    `json:"success_url,omitempty"`
	CancelURL        string    `json:"cancel_url,omitempty"`
	PartiallyPaidURL string    `json:"partially_paid_url,omitempty"`
	Amount           float64   `json:"amount"`
	Currency         string    `json:"currency"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// New create a subscription plan
// JWT is required for this request
func New(su *SubscriptionArgs) (*Subscription, error) {
	if su == nil {
		return nil, errors.New("nil subscription args")
	}

	d, err := json.Marshal(su)
	if err != nil {
		return nil, eris.Wrap(err, "subscription args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "subscription")
	}

	s := &core.V2ResponseFormat[*Subscription]{}
	par := &core.SendParams{
		RouteName: "subscription-create",
		Into:      &s,
		Body:      strings.NewReader(string(d)),
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return s.Result, nil
}

// NewWithEmail create an email subscription with specific plan ID
// JWT is required for this request
func NewWithEmail(su *EmailSubscriptionArgs) (*recurringPayment.RecurringPayment, error) {
	if su == nil {
		return nil, errors.New("nil subscription email args")
	}

	d, err := json.Marshal(su)
	if err != nil {
		return nil, eris.Wrap(err, "subscription email args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "subscription")
	}

	// Inconsistency on their side: this request allow only one e-mail, but respond with an array of RecurringPayment
	// So will return only the first element of array
	s := &core.V2ResponseFormat[[]*recurringPayment.RecurringPayment]{}
	par := &core.SendParams{
		RouteName: "subscription-create-email",
		Into:      &s,
		JWTToken:  tok,
		Body:      strings.NewReader(string(d)),
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return s.Result[0], nil
}

// Update update a subscription plan
// JWT is required for this request
func Update(subscriptionPlanID string, su *SubscriptionArgs) (*Subscription, error) {
	if subscriptionPlanID == "" {
		return nil, eris.New("empty subscription plan ID")
	}

	if su == nil {
		return nil, errors.New("nil subscription args")
	}

	d, err := json.Marshal(su)
	if err != nil {
		return nil, eris.Wrap(err, "subscription args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "subscription")
	}

	s := &core.V2ResponseFormat[*Subscription]{}
	par := &core.SendParams{
		RouteName: "subscription-update",
		Into:      &s,
		JWTToken:  tok,
		Path:      subscriptionPlanID,
		Body:      strings.NewReader(string(d)),
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return s.Result, nil
}

// Get return a single subscription plan by ID
func Get(subscriptionPlanID string) (*Subscription, error) {
	if subscriptionPlanID == "" {
		return nil, eris.New("empty subscription plan ID")
	}

	st := &core.V2ResponseFormat[*Subscription]{}
	par := &core.SendParams{
		RouteName: "subscription-single",
		Path:      subscriptionPlanID,
		Into:      &st,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return st.Result, nil
}
