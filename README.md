# NOWPayments Go Library

[![Go Reference](https://pkg.go.dev/badge/github.com/matm/go-nowpayments.svg)](https://pkg.go.dev/github.com/honsty/go-nowpayments)
[![Go Report Card](https://goreportcard.com/badge/github.com/matm/go-nowpayments)](https://goreportcard.com/report/github.com/honsty/go-nowpayments)
[![codecov](https://codecov.io/gh/matm/go-nowpayments/branch/main/graph/badge.svg?token=AP16BAZR68)](https://codecov.io/gh/honsty/go-nowpayments)

This repository is originally forked from  repository is [https://codecov.io/gh/matn/go-nowpayments](https://codecov.io/gh/matn/go-nowpayments)
This is an unofficial Go library for the [crypto NOWPayments API](https://documenter.getpostman.com/view/7907941/S1a32n38#84c51632-01ad-49c0-96f8-fb4b5ad2b24a)

Topic|Endpoint|Package.Method|Implemented
---|:---|:---|:---:
[Instant Payments Notifications](https://documenter.getpostman.com/view/7907941/S1a32n38#689df54e-9f43-42b3-bfe8-9bcca0444a6a)|||Yes
||Verify signature|[ipn.VerifyRequestSignature(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/ipn#VerifyRequestSignature)|:heavy_check_mark:
[Subscriptions](https://documenter.getpostman.com/view/7907941/2s93JusNJt#7020882a-50d6-465f-bc9b-ff94909bc179)|||Yes
||Create plan|[subscriptions.New(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/subscriptions#New)|:heavy_check_mark:
||Create e-mail subscription|[subscriptions.NewWithEmail(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/subscriptions#NewWithEmail)|:heavy_check_mark:
||Update plan|[subscriptions.Update(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/pkg/subscriptions#Update)|:heavy_check_mark:
||Get plan|[subscriptions.Get(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/subscriptions#Get)|:heavy_check_mark:
||List plans|[subscriptions.List(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/subscriptions#List)|:heavy_check_mark:
[Recurring payments](https://documenter.getpostman.com/view/7907941/S1a32n38#689df54e-9f43-42b3-bfe8-9bcca0444a6a)|||Yes
||Create|[recurring_payments.New(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/recurring_payments#New)|:heavy_check_mark:
||Get|[recurring_payments.Get(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/recurring_payments#Get)|:heavy_check_mark:
||Delete|[recurring_payments.Delete(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/recurring_payments#Delete)|:heavy_check_mark:
[Billing (sub-partner / Custody)](https://documenter.getpostman.com/view/7907941/2s93JusNJt#2b3f0024-d9de-4b91-9db4-d3655e4eded9)|||Yes
||Deposit with payment|[custody.NewDepositWithPayment(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#NewDepositWithPayment)|:heavy_check_mark:
||Deposit from master account|[custody.NewDepositFroMasterAccount(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#NewDepositFroMasterAccount)|:heavy_check_mark:
||Get payments|[custody.GetPayments(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#GetPayments)|:heavy_check_mark:
||Transfer between users|[custody.NewTransfer(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/pkg/custody#NewTransfer)|:heavy_check_mark:
||Get transfer|[custody.GetTransfer(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#GetTransfer)|:heavy_check_mark:
||List transfers|[custody.ListTransfers(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#ListTransfers)|:heavy_check_mark:
||Create user|[custody.NewUser(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#NewUser)|:heavy_check_mark:
||List users|[custody.ListUsers(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#NewUser)|:heavy_check_mark:
||Get user balance|[custody.GetBalance(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#GetBalance)|:heavy_check_mark:
||Write-off to master account|[custody.NewWriteOffToMaster(...)](https://pkg.go.dev/github.com/honsty/go-nowpayments/custody#NewWriteOffToMaster)|:heavy_check_mark:
[Payments](https://documenter.getpostman.com/view/7907941/S1a32n38#84c51632-01ad-49c0-96f8-fb4b5ad2b24a)|||Yes
||Get estimated price|[payments.EstimatedPrice(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#EstimatedPrice)|:heavy_check_mark:
||Get the minimum payment amount|[payments.MinimumAmount(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#MinimumAmount)|:heavy_check_mark:
||Get payment status|[payments.Status()](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#Status)|:heavy_check_mark:
||Get list of payments|[payments.List(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#List)|:heavy_check_mark:
||Get/Update payment estimate|[payments.RefreshEstimatedPrice(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#RefreshEstimatedPrice)|:heavy_check_mark:
||Create invoice|[payments.NewInvoice(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#NewInvoice)|:heavy_check_mark:
||Create payment|[payments.New(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#New)|:heavy_check_mark:
||Create payment from invoice|[payments.NewFromInvoice(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/payments#NewFromInvoice)|:heavy_check_mark:
[Currencies](https://documenter.getpostman.com/view/7907941/S1a32n38#cb80ccdc-8f7c-426c-89df-1ed2241954a5)|||Yes
||Get available currencies|[currencies.All()](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/currencies#All)|:heavy_check_mark:
||Get available checked currencies|[currencies.Selected()](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/currencies#Selected)|:heavy_check_mark:
[Payouts](https://documenter.getpostman.com/view/7907941/S1a32n38#138ee72b-4c4f-40d0-a565-4a1e907f4d94)|||No
[API status](https://documenter.getpostman.com/view/7907941/S1a32n38#9998079f-dcc8-4e07-9ac7-3d52f0fd733a)|||Yes
||Get API status|[core.Status()](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/core#Status)|:heavy_check_mark:
[Authentication](https://documenter.getpostman.com/view/7907941/S1a32n38#174cd8c5-5973-4be7-9213-05567f8adf27)|||Yes
||Authentication|[core.Authenticate(...)](https://pkg.go.dev/github.com/matm/go-nowpayments/pkg/core#Authenticate)|:heavy_check_mark:

## Installation

```bash
$ go get github.com/honsty/go-nowpayments@v1.0.0
```

## Usage

Just load the config with all the credentials from a file or using a `Reader` then display the NOWPayments' API status and the last 2 payments
made with:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/honsty/go-nowpayments/payments"
)

func main() {
	err := config.Load(strings.NewReader(`
            {
                  "server": "https://api-sandbox.nowpayments.io/v1",
                  "login": "some_email@domain.tld",
                  "password": "some_password",
                  "apiKey": "some_api_key"
            }
      `))

	if err != nil {
		log.Fatal(err)
	}

	core.UseBaseURL(core.BaseURL(config.Server()))
	core.UseClient(core.NewHTTPClient())

	ps, err := payments.List(&payments.ListOption{
		Limit: 2,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Last %d payments: %v\n", limit, ps)
}
```

## CLI Tool

The CLI tool has not been updated and is not maintained in this repository
To use it, you can do it from the original repository [https://codecov.io/gh/matn/go-nowpayments](https://codecov.io/gh/matn/go-nowpayments)

