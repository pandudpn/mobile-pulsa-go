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
	
	priceList, err := pricelist.GetWithContext(context.Background(), pricelist.Prepaid, "all", opts)
	if err != nil {
		log.Println(err)
		return
	}
	
	log.Println(priceList)
}
