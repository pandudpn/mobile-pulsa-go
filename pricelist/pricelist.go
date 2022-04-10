package pricelist

import (
	"context"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
)

// Get gets price list of products
func Get(data *PriceListParam, opts *mobilepulsa.Option) (*mobilepulsa.PriceList, error) {
	data.Sign = opts.Sign("pl")
	data.Username = opts.GetUsername()
	data.Commands = "pricelist-pasca"

	r := &request{
		data: data,
		opts: opts,
	}

	return r.GetPriceList(context.Background())
}

// GetWithContext gets price list of products with context
func GetWithContext(ctx context.Context, data *PriceListParam, opts *mobilepulsa.Option) (*mobilepulsa.PriceList, error) {
	data.Sign = opts.Sign("pl")
	data.Username = opts.GetUsername()
	data.Commands = "pricelist-pasca"

	r := &request{
		data: data,
		opts: opts,
	}

	return r.GetPriceList(ctx)
}
