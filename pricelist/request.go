package pricelist

import (
	"context"
	"net/http"
	
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/utils/validator"
)

const (
	// production
	basePostpaidProduction string = "https://mobilepulsa.net"
	basePrepaidProduction  string = "https://prepaid.iak.id"
	// development
	basePostpaidDevelopment string = "https://testpostpaid.mobilepulsa.net"
	basePrepaidDevelopment  string = "https://prepaid.iak.dev"
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
		priceList *mobilepulsa.PriceList
		header    http.Header
		err       error
		url       string
	)
	err = validator.ValidateRequired(ctx, r.data)
	if err != nil {
		return nil, err
	}
	
	if r.data.Service == Postpaid {
		url = basePostpaidProduction + priceListPostpaid
		if r.opts.GetAccessType() == mobilepulsa.Development {
			url = basePostpaidDevelopment + priceListPostpaid
		}
	} else {
		url = basePrepaidProduction + priceListPrepaid
		if r.opts.GetAccessType() == mobilepulsa.Development {
			url = basePrepaidDevelopment + priceListPrepaid
		}
	}
	r.data.Sign = r.opts.Sign("pl")
	
	err = r.opts.GetAPIRequest().Call(ctx, http.MethodPost, url, header, r.data, &priceList)
	if err != nil {
		return nil, err
	}
	
	return priceList, nil
}
