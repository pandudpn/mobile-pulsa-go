package topup_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/topup"
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

	result.(*mobilepulsa.TopUp).Data = mobilepulsa.DataTopUp{
		RefID:       "123",
		Status:      0,
		ProductCode: "htelkomsel1000",
		RC:          "00",
		Message:     "",
		Price:       1900,
		TrID:        23,
	}

	return nil
}

func TestCreatePayment(t *testing.T) {
	apiRequestMockObj := new(apiRequestMock)
	opts := initTesting(apiRequestMockObj)

	testCases := []struct {
		name           string
		context        context.Context
		url            string
		data           *topup.TopUpParam
		expectedResult *mobilepulsa.TopUp
		expectedErr    error
	}{
		{
			name:    "success create payment",
			context: nil,
			url:     "https://prepaid.iak.dev/api/top-up",
			data: &topup.TopUpParam{
				ProductCode: "htelkomsel1000",
				RefID:       "123",
				CustomerID:  "0822222222",
			},
			expectedResult: &mobilepulsa.TopUp{
				Data: mobilepulsa.DataTopUp{
					RefID:       "123",
					Status:      0,
					ProductCode: "htelkomsel1000",
					RC:          "00",
					Message:     "",
					Price:       1900,
					TrID:        23,
				},
			},
			expectedErr: nil,
		},
		{
			name:    "success create payment with context",
			context: context.Background(),
			url:     "https://prepaid.iak.dev/api/top-up",
			data: &topup.TopUpParam{
				ProductCode: "htelkomsel1000",
				RefID:       "123",
				CustomerID:  "0822222222",
			},
			expectedResult: &mobilepulsa.TopUp{
				Data: mobilepulsa.DataTopUp{
					RefID:       "123",
					Status:      0,
					ProductCode: "htelkomsel1000",
					RC:          "00",
					Message:     "",
					Price:       1900,
					TrID:        23,
				},
			},
			expectedErr: nil,
		},
		{
			name:    "error given missing required fields",
			context: nil,
			url:     "https://prepaid.iak.dev/api/top-up",
			data: &topup.TopUpParam{
				ProductCode: "",
				RefID:       "123",
				CustomerID:  "0822222222",
			},
			expectedResult: nil,
			expectedErr:    errors.New("missing required fields: ProductCode"),
		},
		{
			name:    "error given missing required fields",
			context: nil,
			url:     "",
			data: &topup.TopUpParam{
				ProductCode: "",
				RefID:       "123",
				CustomerID:  "0822222222",
			},
			expectedResult: nil,
			expectedErr:    errors.New("missing required fields: ProductCode"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				header http.Header
				resp   *mobilepulsa.TopUp
				err    error
			)
			apiRequestMockObj.On(
				"Call",
				context.Background(),
				http.MethodPost,
				tc.url,
				header,
				tc.data,
				&mobilepulsa.TopUp{},
			)

			if tc.context == nil {
				resp, err = topup.CreatePayment(tc.data, opts)
			} else {
				resp, err = topup.CreatePaymentWithContext(tc.context, tc.data, opts)
			}

			assert.Equal(t, tc.expectedResult, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
