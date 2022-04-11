package main

import (
	"log"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/topup"
)

func main() {
	opts := mobilepulsa.NewOption()
	opts.SetAPIKey("api-key")
	opts.SetUsername("username")
	opts.SetAccessDevelopment()

	data := &topup.TopUpParam{
		RefID:       "ref-id-123",
		CustomerID:  "082222222",
		ProductCode: "htelkomsel1000",
	}

	topup, err := topup.CreatePayment(data, opts)
	if err != nil {
		log.Println("error create payment topup", err)
		return
	}

	log.Println(topup)
}
