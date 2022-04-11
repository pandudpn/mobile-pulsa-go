package topup

import (
	"context"
	"net/http"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/utils/validator"
)

const (
	// endpoint
	topUpPrepaid string = "/api/top-up"
)

type requestTopUp struct {
	data *TopUpParam
	opts *mobilepulsa.Option
}

func (r *requestTopUp) CreatePayment(ctx context.Context) (*mobilepulsa.TopUp, error) {
	var (
		topUp  mobilepulsa.TopUp
		header http.Header
		err    error
		url    = mobilepulsa.BasePrepaidProduction + topUpPrepaid
	)
	err = validator.ValidateRequired(ctx, r.data)
	if err != nil {
		return nil, err
	}

	if r.opts.GetAccessType() == mobilepulsa.Development {
		url = mobilepulsa.BasePrepaidDevelopment + topUpPrepaid
	}

	err = r.opts.GetAPIRequest().Call(ctx, http.MethodPost, url, header, r.data, &topUp)
	if err != nil {
		return nil, err
	}

	return &topUp, nil
}
