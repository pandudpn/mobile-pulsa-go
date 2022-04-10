package main

import (
	"context"
	"log"
	
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/pricelist"
)

func main() {
	opts := mobilepulsa.NewOption()
	opts.SetAPIKey("5576252b0081a24e")
	opts.SetUsername("083875181609")
	opts.SetAccessDevelopment()
	
	priceList, err := pricelist.GetWithContext(context.Background(), pricelist.Prepaid, "all", opts)
	if err != nil {
		log.Println(err)
		return
	}
	
	log.Println(priceList)
}
