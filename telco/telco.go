package telco

import (
	"context"
	"strconv"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
)

// CreateInquiry will create inquiry before payment for Postpaid Product
// and will return a data for Payment
func CreateInquiry(data *InquiryParam, opts *mobilepulsa.Option) (*mobilepulsa.Inquiry, error) {
	data.Commands = "inq-pasca"
	data.Username = opts.GetUsername()
	data.Sign = opts.Sign(data.RefID)

	r := &requestTelco{
		data: data,
		opts: opts,
	}

	return r.CreateInquiry(context.Background())
}

// CreateInquiryWithContext will create inquiry before payment for Postpaid Product
// and will return a data for Payment with context
func CreateInquiryWithContext(ctx context.Context, data *InquiryParam, opts *mobilepulsa.Option) (*mobilepulsa.Inquiry, error) {
	data.Commands = "inq-pasca"
	data.Username = opts.GetUsername()
	data.Sign = opts.Sign(data.RefID)

	r := &requestTelco{
		data: data,
		opts: opts,
	}

	return r.CreateInquiry(ctx)
}

// CreatePayment will create payment for Postpaid Product
func CreatePayment(data *PaymentParam, opts *mobilepulsa.Option) (*mobilepulsa.Payment, error) {
	trId := strconv.Itoa(data.TrID)

	data.Commands = "inq-pasca"
	data.Username = opts.GetUsername()
	data.Sign = opts.Sign(trId)

	r := &requestTelco{
		data: data,
		opts: opts,
	}

	return r.CreatePayment(context.Background())
}

// CreatePaymentWithContext will create payment for Postpaid Product with context
func CreatePaymentWithContext(ctx context.Context, data *PaymentParam, opts *mobilepulsa.Option) (*mobilepulsa.Payment, error) {
	trId := strconv.Itoa(data.TrID)

	data.Commands = "pay-pasca"
	data.Username = opts.GetUsername()
	data.Sign = opts.Sign(trId)

	r := &requestTelco{
		data: data,
		opts: opts,
	}

	return r.CreatePayment(ctx)
}
