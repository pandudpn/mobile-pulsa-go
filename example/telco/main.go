package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/telco"
)

func main() {
	opts := mobilepulsa.NewOption()
	opts.SetAPIKey(os.Getenv("IAK_API_KEY"))
	opts.SetUsername(os.Getenv("IAK_USERNAME"))
	opts.SetAccessDevelopment()

	data := &telco.InquiryParam{
		RefID: "ref-123456762",
		Code:  "HPTHREE",
		HP:    "08991234501",
	}

	inquiry, err := telco.CreateInquiryWithContext(context.Background(), data, opts)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(inquiry)

	paymentData := &telco.PaymentParam{
		TrID: inquiry.Data.TrID,
	}

	payment, err := telco.CreatePaymentWithContext(context.Background(), paymentData, opts)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(payment)
}
