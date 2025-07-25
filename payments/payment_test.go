package payments

import (
	"errors"
	"net/http"
	"testing"

	"github.com/honsty/go-nowpayments/core"
	"github.com/honsty/go-nowpayments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		pa    *PaymentArgs
		init  func(*mocks.HTTPClient)
		after func(*Payment[string], error)
	}{
		{"nil args", nil, nil,
			func(p *Payment[string], err error) {
				assert.Nil(p)
				assert.Error(err)
			},
		},
		{"api error", &PaymentArgs{PurchaseID: "1234"},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(p *Payment[string], err error) {
				assert.Nil(p)
				assert.Error(err)
				assert.Equal("payment-create: network error", err.Error())
			},
		},
		{"valid args", &PaymentArgs{
			PurchaseID:    "1234",
			PaymentAmount: PaymentAmount{PriceAmount: 10.0},
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
				t.Logf("%+v", p)
			},
		},
		{"pay_amount as a string", &PaymentArgs{
			PurchaseID:    "1234",
			PaymentAmount: PaymentAmount{PriceAmount: 10.0},
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234","pay_amount":3.5}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
				assert.Equal(3.5, p.PayAmount)
			},
		},
		{"pay_amount as a float", &PaymentArgs{
			PurchaseID:    "1234",
			PaymentAmount: PaymentAmount{PriceAmount: 10.0},
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234","pay_amount":4.2}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
				assert.Equal(4.2, p.PayAmount)
			},
		},
		{"pay_amount as an integer, who knows...", &PaymentArgs{
			PurchaseID:    "1234",
			PaymentAmount: PaymentAmount{PriceAmount: 10.0},
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234","pay_amount":100}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
			},
		},
		{"missing pay_amount value", &PaymentArgs{
			PurchaseID:    "1234",
			PaymentAmount: PaymentAmount{PriceAmount: 10.0},
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
			},
		},
		{"route check", &PaymentArgs{},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Run(func(r *http.Request) {
					assert.Equal("/v1/payment", r.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewHTTPClient(t)
			core.UseClient(c)
			if tt.init != nil {
				tt.init(c)
			}
			got, err := New(tt.pa)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}

func TestNewFromInvoice(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		ipa   *InvoicePaymentArgs
		init  func(*mocks.HTTPClient)
		after func(*Payment[string], error)
	}{
		{"route check", &InvoicePaymentArgs{},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Run(func(r *http.Request) {
					assert.Equal("/v1/invoice-payment", r.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
			},
		},
		{"valid args", &InvoicePaymentArgs{
			InvoiceID:  "55",
			PurchaseID: "1234",
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment[string], err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
			},
		},
		{"nil args", nil, nil,
			func(p *Payment[string], err error) {
				assert.Nil(p)
				assert.Error(err)
			},
		},
		{"api error", &InvoicePaymentArgs{InvoiceID: "1234"},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(p *Payment[string], err error) {
				assert.Nil(p)
				assert.Error(err)
				assert.Equal("invoice-payment: network error", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewHTTPClient(t)
			core.UseClient(c)
			if tt.init != nil {
				tt.init(c)
			}
			got, err := NewFromInvoice(tt.ipa)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
