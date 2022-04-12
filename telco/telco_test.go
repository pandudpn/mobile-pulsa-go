package telco_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/pandudpn/mobile-pulsa-go/telco"
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

	_, ok := body.(*telco.InquiryParam)
	if ok {
		result.(*mobilepulsa.Inquiry).Data = mobilepulsa.DataInquiry{
			TrID:  123,
			RefID: "ref-123",
			Code:  "HPTHREE",
			HP:    "08991234501",
		}
	}

	_, ok = body.(*telco.PaymentParam)
	if ok {
		result.(*mobilepulsa.Payment).Data = mobilepulsa.DataPayment{
			TrID:  123,
			RefID: "ref-123",
			Code:  "HPTHREE",
			HP:    "08991234501",
		}
	}

	return nil
}

func TestCreateInquiry(t *testing.T) {
	apiRequestMockObj := new(apiRequestMock)
	opts := initTesting(apiRequestMockObj)

	testCases := []struct {
		name           string
		context        context.Context
		url            string
		data           *telco.InquiryParam
		expectedResult *mobilepulsa.Inquiry
		expectedErr    error
	}{
		{
			name:    "success create inquiry",
			context: nil,
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &telco.InquiryParam{
				HP:    "08991234501",
				Code:  "HPTHREE",
				RefID: "ref-123",
			},
			expectedResult: &mobilepulsa.Inquiry{
				Data: mobilepulsa.DataInquiry{
					TrID:  123,
					RefID: "ref-123",
					Code:  "HPTHREE",
					HP:    "08991234501",
				},
			},
			expectedErr: nil,
		},
		{
			name:    "success create inquiry with context",
			context: context.Background(),
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &telco.InquiryParam{
				HP:    "08991234501",
				Code:  "HPTHREE",
				RefID: "ref-123",
			},
			expectedResult: &mobilepulsa.Inquiry{
				Data: mobilepulsa.DataInquiry{
					TrID:  123,
					RefID: "ref-123",
					Code:  "HPTHREE",
					HP:    "08991234501",
				},
			},
			expectedErr: nil,
		},
		{
			name:    "error given missing required fields",
			context: nil,
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &telco.InquiryParam{
				HP:    "",
				Code:  "HPTHREE",
				RefID: "ref-123",
			},
			expectedResult: nil,
			expectedErr:    errors.New("missing required fields: HP"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				header http.Header
				resp   *mobilepulsa.Inquiry
				err    error
			)
			apiRequestMockObj.On(
				"Call",
				context.Background(),
				http.MethodPost,
				tc.url,
				header,
				tc.data,
				&mobilepulsa.Inquiry{},
			)

			if tc.context == nil {
				resp, err = telco.CreateInquiry(tc.data, opts)
			} else {
				resp, err = telco.CreateInquiryWithContext(tc.context, tc.data, opts)
			}

			assert.Equal(t, tc.expectedResult, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestCreatePayment(t *testing.T) {
	apiRequestMockObj := new(apiRequestMock)
	opts := initTesting(apiRequestMockObj)

	testCases := []struct {
		name           string
		context        context.Context
		url            string
		data           *telco.PaymentParam
		expectedResult *mobilepulsa.Payment
		expectedErr    error
	}{
		{
			name:    "success create inquiry",
			context: nil,
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &telco.PaymentParam{
				TrID: 123,
			},
			expectedResult: &mobilepulsa.Payment{
				Data: mobilepulsa.DataPayment{
					TrID:  123,
					RefID: "ref-123",
					Code:  "HPTHREE",
					HP:    "08991234501",
				},
			},
			expectedErr: nil,
		},
		{
			name:    "success create inquiry with context",
			context: context.Background(),
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &telco.PaymentParam{
				TrID: 123,
			},
			expectedResult: &mobilepulsa.Payment{
				Data: mobilepulsa.DataPayment{
					TrID:  123,
					RefID: "ref-123",
					Code:  "HPTHREE",
					HP:    "08991234501",
				},
			},
			expectedErr: nil,
		},
		{
			name:    "error given missing required fields",
			context: nil,
			url:     "https://testpostpaid.mobilepulsa.net/api/v1/bill/check",
			data: &telco.PaymentParam{
				TrID: 0,
			},
			expectedResult: nil,
			expectedErr:    errors.New("missing required fields: TrID"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				header http.Header
				resp   *mobilepulsa.Payment
				err    error
			)
			apiRequestMockObj.On(
				"Call",
				context.Background(),
				http.MethodPost,
				tc.url,
				header,
				tc.data,
				&mobilepulsa.Payment{},
			)

			if tc.context == nil {
				resp, err = telco.CreatePayment(tc.data, opts)
			} else {
				resp, err = telco.CreatePaymentWithContext(tc.context, tc.data, opts)
			}

			assert.Equal(t, tc.expectedResult, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
