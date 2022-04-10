package mobilepulsa_test

import (
	"context"
	"net/http"
	"testing"

	mobilepulsa "github.com/pandudpn/mobile-pulsa-go"
	"github.com/stretchr/testify/mock"
)

type apiRequestMock struct {
	mock.Mock
}

func (a *apiRequestMock) Call(ctx context.Context, httpMethod, url string, header http.Header, body interface{}, result interface{}) error {
	a.Called(ctx, httpMethod, url, header, body, result)

	return nil
}

func TestNewOption(t *testing.T) {
	testCases := []struct {
		name        string
		username    string
		apikey      string
		development bool
		httpClient  *http.Client
	}{
		{
			name:        "success valid",
			username:    "test",
			apikey:      "test-api-key",
			development: true,
			httpClient:  &http.Client{},
		},
		{
			name:        "get error username is nil",
			username:    "",
			apikey:      "test-api-key",
			development: true,
			httpClient:  &http.Client{},
		},
		{
			name:        "get error api-key is nil",
			username:    "test",
			apikey:      "",
			development: false,
			httpClient:  &http.Client{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			apiRequestMockObj := new(apiRequestMock)

			opts := mobilepulsa.NewOption()
			opts.SetUsername(tc.username)
			opts.SetAPIKey(tc.apikey)
			opts.SetHTTPClient(tc.httpClient)
			opts.SetAPIRequest(apiRequestMockObj)
			opts.Valid()

			if tc.development {
				opts.SetAccessDevelopment()
			} else {
				opts.SetAccessProduction()
			}
		})
	}
}
