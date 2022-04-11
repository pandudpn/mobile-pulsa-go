package topup

import (
	"context"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
)

// CreatePayment create payment top up for prepaid service
func CreatePayment(data *TopUpParam, opts *mobilepulsa.Option) (*mobilepulsa.TopUp, error) {
	data.Username = opts.GetUsername()
	data.Sign = opts.Sign(data.RefID)

	r := &requestTopUp{
		data: data,
		opts: opts,
	}

	return r.CreatePayment(context.Background())
}

// CreatePaymentWithContext create payment top up for prepaid service with context
func CreatePaymentWithContext(ctx context.Context, data *TopUpParam, opts *mobilepulsa.Option) (*mobilepulsa.TopUp, error) {
	data.Username = opts.GetUsername()
	data.Sign = opts.Sign(data.RefID)

	r := &requestTopUp{
		data: data,
		opts: opts,
	}

	return r.CreatePayment(ctx)
}
