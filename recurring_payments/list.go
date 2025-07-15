package recurring_payments

import (
	"fmt"
	"net/url"

	"github.com/honsty/go-nowpayments/core"
)

type ListOption struct {
	Limit              int
	Offset             int
	IsActive           *bool
	Status             *string
	SubscriptionPlanID *int64
}

// List returns a list of all recurring payments, depending on the supplied options (which can be nil)
func List(o *ListOption) ([]*RecurringPayment, error) {
	u := url.Values{}

	if o != nil {
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		if o.Offset != 0 {
			u.Set("offset", fmt.Sprintf("%d", o.Offset))
		}
		if o.IsActive != nil {
			u.Set("is_active", fmt.Sprintf("%d", o.IsActive))
		}
		if o.Status != nil {
			u.Set("status", fmt.Sprintf("%d", o.Status))
		}
		if o.SubscriptionPlanID != nil {
			u.Set("subscription_plan_id", fmt.Sprintf("%d", o.SubscriptionPlanID))
		}
	}

	rpl := &core.V2ResponseFormat[[]*RecurringPayment]{}
	par := &core.SendParams{
		RouteName: "recurring-payment-list",
		Into:      rpl,
		Values:    u,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return rpl.Result, nil
}
