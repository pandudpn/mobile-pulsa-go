package mobilepulsa_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/pricelist"
)

func initTesting() *mobilepulsa.Option {
	opts := mobilepulsa.NewOption()
	opts.SetUsername("abc")
	opts.SetAPIKey("abc")
	
	return opts
}

func TestAPIRequestImplementation_Call(t *testing.T) {
	opts := initTesting()
	
	testCases := []struct {
		name           string
		context        context.Context
		url            string
		data           interface{}
		expectedResult *mobilepulsa.PriceList
		expectedErr    error
	}{
		{
			name:    "success get price list products prepaid",
			context: context.Background(),
			url:     "https://prepaid.iak.dev/api/pricelist",
			data: &pricelist.PriceListParam{
				Service: pricelist.Prepaid,
				Status:  "all",
			},
			expectedResult: &mobilepulsa.PriceList{
				Data: mobilepulsa.DataPricelist{
					PriceList: []mobilepulsa.DataPrepaid{
						{
							ProductCode:        "htelkomsel1000",
							ProductPrice:       1900,
							ProductNominal:     "1000",
							ProductDescription: "Telkomsel",
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:    "success get price list products postpaid",
			context: context.Background(),
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &pricelist.PriceListParam{
				Service: pricelist.Postpaid,
				Status:  "all",
			},
			expectedResult: &mobilepulsa.PriceList{
				Data: mobilepulsa.DataPricelist{
					Pasca: []mobilepulsa.DataPostpaid{
						{
							Code:   "BPJS",
							Name:   "BPJS Kesehatan",
							Fee:    2500,
							Komisi: 1150,
							Status: 1,
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:    "error given missing required fields",
			context: nil,
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &pricelist.PriceListParam{
				Service: pricelist.Postpaid,
			},
			expectedResult: nil,
			expectedErr:    errors.New("missing required fields: Status"),
		},
		{
			name:           "error given json unmarshal",
			context:        nil,
			url:            "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data:           "<",
			expectedResult: nil,
			expectedErr:    errors.New("invalid character '<' looking for beginning of value"),
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				header http.Header
			)
			
			opts.GetAPIRequest().Call(context.Background(), http.MethodPost, tc.url, header, tc.data, tc.expectedResult)
		})
	}
}
