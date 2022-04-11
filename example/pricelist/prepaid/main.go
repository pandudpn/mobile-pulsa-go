package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/pricelist"
)

func main() {
	opts := mobilepulsa.NewOption()
	opts.SetAPIKey(os.Getenv("IAK_API_KEY"))
	opts.SetUsername(os.Getenv("IAK_USERNAME"))
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
