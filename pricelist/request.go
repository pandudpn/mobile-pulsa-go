package pricelist

import (
	"context"
	"net/http"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/utils/validator"
)

const (
	// endpoint
	priceListPostpaid string = "/api/v1/bill/check"
	priceListPrepaid  string = "/api/pricelist"
)

type request struct {
	opts *mobilepulsa.Option
	data *PriceListParam
}

func (r *request) GetPriceList(ctx context.Context) (*mobilepulsa.PriceList, error) {
	var (
		priceList mobilepulsa.PriceList
		header    http.Header
		err       error
		url       string
	)
	err = validator.ValidateRequired(ctx, r.data)
	if err != nil {
		return nil, err
	}

	if r.data.Service == Postpaid {
		url = mobilepulsa.BasePostpaidProduction + priceListPostpaid
		if r.opts.GetAccessType() == mobilepulsa.Development {
			url = mobilepulsa.BasePostpaidDevelopment + priceListPostpaid
		}
	} else {
		url = mobilepulsa.BasePrepaidProduction + priceListPrepaid
		if r.opts.GetAccessType() == mobilepulsa.Development {
			url = mobilepulsa.BasePrepaidDevelopment + priceListPrepaid
		}
	}

	err = r.opts.GetAPIRequest().Call(ctx, http.MethodPost, url, header, r.data, &priceList)
	if err != nil {
		return nil, err
	}

	return &priceList, nil
}
