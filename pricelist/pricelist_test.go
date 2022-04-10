package pricelist_test

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
	
	"github.com/go-playground/assert/v2"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/pricelist"
	"github.com/stretchr/testify/mock"
)

type apiRequestMock struct {
	mock.Mock
}

func initTesting(apiRequestMockObj mobilepulsa.APIRequest) *mobilepulsa.Option {
	opts := mobilepulsa.NewOption()
	opts.SetUsername("abc")
	opts.SetAPIKey("abc")
	opts.SetAPIRequest(apiRequestMockObj)
	
	return opts
}

func (a *apiRequestMock) Call(ctx context.Context, httpMethod, url string, header http.Header, body interface{}, result interface{}) error {
	a.Called(ctx, httpMethod, url, header, body, result)
	
	if strings.Contains(url, "prepaid") {
		result.(*mobilepulsa.PriceList).Data = mobilepulsa.DataPricelist{
			PriceList: []mobilepulsa.DataPrepaid{
				{
					ProductCode:        "htelkomsel1000",
					ProductPrice:       1900,
					ProductNominal:     "1000",
					ProductDescription: "Telkomsel",
				},
			},
		}
	} else {
		result.(*mobilepulsa.PriceList).Data = mobilepulsa.DataPricelist{
			Pasca: []mobilepulsa.DataPostpaid{
				{
					Code:   "BPJS",
					Name:   "BPJS Kesehatan",
					Fee:    2500,
					Komisi: 1150,
					Status: 1,
				},
			},
		}
	}
	
	return nil
}

func TestGet(t *testing.T) {
	apiRequestMockObj := new(apiRequestMock)
	opts := initTesting(apiRequestMockObj)
	
	testCases := []struct {
		name           string
		context        context.Context
		url            string
		data           *pricelist.PriceListParam
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
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				header http.Header
				resp   *mobilepulsa.PriceList
				err    error
			)
			apiRequestMockObj.On(
				"Call",
				context.Background(),
				http.MethodPost,
				tc.url,
				header,
				tc.data,
				&mobilepulsa.PriceList{},
			)
			
			if tc.context == nil {
				resp, err = pricelist.Get(tc.data, opts)
			} else {
				resp, err = pricelist.GetWithContext(tc.context, tc.data, opts)
			}
			
			assert.Equal(t, tc.expectedResult, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
