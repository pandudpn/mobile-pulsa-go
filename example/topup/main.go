package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/topup"
)

func main() {
	opts := mobilepulsa.NewOption()
	opts.SetAPIKey(os.Getenv("IAK_API_KEY"))
	opts.SetUsername(os.Getenv("IAK_USERNAME"))
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
