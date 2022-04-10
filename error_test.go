package mobilepulsa_test

import (
	"errors"
	"testing"
	
	"github.com/go-playground/assert/v2"
	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
)

func TestErrorHttp(t *testing.T) {
	testCases := []struct {
		name           string
		resBody        string
		expectedResult mobilepulsa.ErrorCode
	}{
		{
			name:           "success handle error postpaid and get ErrInvoicePaid",
			resBody:        `{"data": {"response_code": "01", "message": ""}}`,
			expectedResult: mobilepulsa.ErrInvoicePaid,
		},
		{
			name:           "success handle error prepaid and get ErrUsername",
			resBody:        `{"data": {"rc": "208", "message": "INVALID DATA", "status": 2}}`,
			expectedResult: mobilepulsa.ErrUsername,
		},
		{
			name:           "error unmarshal",
			resBody:        "<",
			expectedResult: errors.New("invalid character '<' looking for beginning of value"),
		},
		{
			name:           "success handle error un-mapped key and get ErrParseFailed",
			resBody:        `{"data": {"rc": "210", "message": "INVALID DATA", "status": 2}}`,
			expectedResult: mobilepulsa.ErrParseFailed,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := mobilepulsa.ErrorHttp([]byte(tc.resBody))
			
			assert.Equal(t, tc.expectedResult.Error(), err.Error())
		})
	}
}
