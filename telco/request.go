package telco

import (
	"context"
	"net/http"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/utils/validator"
)

const postpaid = "/api/v1/bill/check"

type requestTelco struct {
	data interface{}
	opts *mobilepulsa.Option
}

// CreateInquiry will request to mobilepulsa.net for product Telco Postpaid
func (r *requestTelco) CreateInquiry(ctx context.Context) (*mobilepulsa.Inquiry, error) {
	var (
		inq    mobilepulsa.Inquiry
		header http.Header
		err    error
		url    = mobilepulsa.BasePostpaidProduction
	)

	err = validator.ValidateRequired(ctx, r.data)
	if err != nil {
		return nil, err
	}

	if r.opts.GetAccessType() == mobilepulsa.Development {
		url = mobilepulsa.BasePostpaidDevelopment
	}
	url += postpaid

	err = r.opts.GetAPIRequest().Call(ctx, http.MethodPost, url, header, r.data, &inq)
	if err != nil {
		return nil, err
	}

	return &inq, nil
}

// CreatePayment will request to mobilepulsa.net for product Telco Postpaid
func (r *requestTelco) CreatePayment(ctx context.Context) (*mobilepulsa.Payment, error) {
	var (
		payment mobilepulsa.Payment
		header  http.Header
		err     error
		url     = mobilepulsa.BasePostpaidProduction
	)

	err = validator.ValidateRequired(ctx, r.data)
	if err != nil {
		return nil, err
	}

	if r.opts.GetAccessType() == mobilepulsa.Development {
		url = mobilepulsa.BasePostpaidDevelopment
	}
	url += postpaid

	err = r.opts.GetAPIRequest().Call(ctx, http.MethodPost, url, header, r.data, &payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
