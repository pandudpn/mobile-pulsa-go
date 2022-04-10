package pricelist

import (
	"context"
	
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
)

// Get gets price list of products
func Get(service Service, status string, opts *mobilepulsa.Option) (*mobilepulsa.PriceList, error) {
	data := &priceListParam{
		Service:  service,
		Status:   status,
		Username: opts.GetUsername(),
		Sign:     opts.Sign("pl"),
		Commands: "pricelist-pasca",
	}
	
	r := &request{
		data: data,
		opts: opts,
	}
	
	return r.GetPriceList(context.Background())
}

// GetWithContext gets price list of products with context
func GetWithContext(ctx context.Context, service Service, status string, opts *mobilepulsa.Option) (*mobilepulsa.PriceList, error) {
	data := &priceListParam{
		Service:  service,
		Status:   status,
		Username: opts.GetUsername(),
		Sign:     opts.Sign("pl"),
		Commands: "pricelist-pasca",
	}
	
	r := &request{
		data: data,
		opts: opts,
	}
	
	return r.GetPriceList(ctx)
}
