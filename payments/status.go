package payments

import (
	"github.com/honsty/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// PaymentStatus is the actual information about a payment
type PaymentStatus struct {
	ID             int64   `json:"payment_id"`
	InvoiceID      int64   `json:"invoice_id"`
	Status         string  `json:"payment_status"`
	PayAddress     string  `json:"pay_address"`
	PayinExtraID   string  `json:"payin_extra_id"`
	PriceAmount    float64 `json:"price_amount"`
	PriceCurrency  string  `json:"price_currency"`
	PayAmount      float64 `json:"pay_amount"`
	ActuallyPaid   float64 `json:"actually_paid"`
	PayCurrency    string  `json:"pay_currency"`
	OrderID        string  `json:"order_id"`
	PurchaseID     int64   `json:"purchase_id"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	BurningPurcent string  `json:"burning_percent"`
	Type           string  `json:"type"`
}

// Status gets the actual information about the payment. You need to provide the payment ID
func Status(paymentID string) (*PaymentStatus, error) {
	if paymentID == "" {
		return nil, eris.New("empty payment ID")
	}

	st := &PaymentStatus{}
	par := &core.SendParams{
		RouteName: "payment-status",
		Path:      paymentID,
		Into:      &st,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return st, nil
}
