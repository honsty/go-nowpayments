package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/honsty/go-nowpayments/config"
	"github.com/honsty/go-nowpayments/core"
	"github.com/honsty/go-nowpayments/currencies"
	"github.com/honsty/go-nowpayments/payments"
)

func main() {
	cfgFile := flag.String("f", "", "JSON config file to use")
	paymentID := flag.String("p", "", "status of payment ID")
	newPayment := flag.Bool("n", false, "new payment")
	payAmount := flag.Float64("a", 2.0, "pay amount for new payment/invoice")
	payCurrency := flag.String("pc", "xmr", "crypto currency to pay in")
	listPayments := flag.Bool("l", false, "list all payments")
	debug := flag.Bool("debug", false, "turn debugging on")
	showCurrencies := flag.Bool("c", false, "show list of selected currencies")
	newInvoice := flag.Bool("i", false, "new invoice")
	newPaymentFromInvoice := flag.String("pi", "", "new payment from invoice ID")
	pcase := flag.String("case", "success", "payment's case (sandbox only)")

	flag.Parse()

	if *cfgFile == "" {
		log.Fatal("please specify a JSON config file with -f")
	}
	f, err := os.Open(*cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = config.LoadFromFile(f)
	if err != nil {
		log.Fatal(err)
	}
	core.UseBaseURL(core.BaseURL(config.Server()))
	core.UseClient(core.NewHTTPClient())

	if *debug {
		core.WithDebug(true)
		fmt.Fprintln(os.Stderr, "Debug:", *debug)
	}

	if *paymentID != "" {
		ps, err := payments.Status(*paymentID)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(ps)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
		return
	}

	fmt.Fprintln(os.Stderr, "Sandbox:", config.Server() == core.SandBoxBaseURL)
	st, err := core.Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, "API Status:", st)

	if *showCurrencies {
		cs, err := currencies.Selected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(os.Stderr, "%d selected (checked) crypto currencies: %v\n", len(cs), cs)
		return
	}

	if *listPayments {
		limit := 5
		fmt.Fprintf(os.Stderr, "Showing last %d payments only:\n", limit)
		all, err := payments.List(&payments.ListOption{
			Limit: limit,
		})
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(all)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
		return
	}

	if *newPayment {
		pa := &payments.PaymentArgs{
			PaymentAmount: payments.PaymentAmount{
				PriceAmount:      *payAmount,
				PriceCurrency:    "eur",
				PayCurrency:      *payCurrency,
				OrderID:          "tool 1",
				OrderDescription: "Some useful tool",
			},
		}
		if config.Server() == core.SandBoxBaseURL {
			pa.Case = *pcase
		}
		fmt.Fprintf(os.Stderr, "Creating a %.2f payment ...\n", pa.PriceAmount)
		pay, err := payments.New(pa)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(pay)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
		return
	}

	if *newInvoice {
		pa := &payments.InvoiceArgs{
			PaymentAmount: payments.PaymentAmount{
				PriceAmount:      *payAmount,
				PriceCurrency:    "eur",
				PayCurrency:      "xmr",
				OrderID:          "tool 1",
				OrderDescription: "Some useful tool",
			},
			CancelURL:  "http://mycancel",
			SuccessURL: "http://mysuccess",
		}
		fmt.Fprintf(os.Stderr, "Creating a %.2f invoice ...\n", pa.PriceAmount)
		pay, err := payments.NewInvoice(pa)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(pay)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
		return
	}

	if *newPaymentFromInvoice != "" {
		pa := &payments.InvoicePaymentArgs{
			InvoiceID:   *newPaymentFromInvoice,
			PayCurrency: "xmr",
		}
		fmt.Fprintf(os.Stderr, "Creating a payment from invoice %q...\n", pa.InvoiceID)
		pay, err := payments.NewFromInvoice(pa)
		if err != nil {
			log.Fatal(err)
		}
		d, err := json.Marshal(pay)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(d))
		return
	}
}
