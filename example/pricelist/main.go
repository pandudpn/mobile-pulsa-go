package main

import (
	"context"
	"log"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/pricelist"
)

func main() {
	opts := mobilepulsa.NewOption()
	opts.SetAPIKey("api-key")
	opts.SetUsername("username")
	opts.SetAccessDevelopment()

	data := &pricelist.PriceListParam{
		Status:  "all",
		Service: pricelist.Prepaid,
	}

	priceList, err := pricelist.GetWithContext(context.Background(), data, opts)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(priceList)
}
