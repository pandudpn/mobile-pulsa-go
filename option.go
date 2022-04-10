package mobilepulsa

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"reflect"
)

type AccessType string

const (
	Production  AccessType = "production"
	Development AccessType = "development"
)

// Option is the wrap of the parameters needed for this API Call
type Option struct {
	apiKey     string
	userName   string
	accessType AccessType
	apiCall    APIRequest
}

// NewOption creates new Option parameter's for API Call
func NewOption() *Option {
	httpClient := &http.Client{}

	return &Option{
		accessType: Development,
		apiCall: &APIRequestImplementation{
			HTTPClient: httpClient,
		},
	}
}

// SetAPIKey will set api-key into parameter apiKey
func (o *Option) SetAPIKey(apiKey string) {
	o.apiKey = apiKey
}

// SetUsername will set username into parameter username
func (o *Option) SetUsername(username string) {
	o.userName = username
}

// GetUsername will return a username of parameter
func (o *Option) GetUsername() string {
	return o.userName
}

// SetAccessProduction will set environment production into parameter accessType
func (o *Option) SetAccessProduction() {
	o.accessType = Production
}

// SetAccessDevelopment will set environment development into parameter accessType
func (o *Option) SetAccessDevelopment() {
	o.accessType = Development
}

// GetAccessType will set environment production into parameter accessType
func (o *Option) GetAccessType() AccessType {
	return o.accessType
}

// SetHTTPClient will set http client into parameter API Call
func (o *Option) SetHTTPClient(httpClient *http.Client) {
	o.apiCall = &APIRequestImplementation{
		HTTPClient: httpClient,
	}
}

// SetAPIRequest will set standard API Request
func (o *Option) SetAPIRequest(apiCall APIRequest) {
	o.apiCall = apiCall
}

// GetAPIRequest will get an instance of APIRequest
func (o *Option) GetAPIRequest() APIRequest {
	return o.apiCall
}

// Valid is for checking required parameter's for API Call
func (o *Option) Valid() error {
	if o.apiKey == "" || reflect.ValueOf(o.apiKey).IsZero() {
		return ErrAPIKeyNil
	}

	if o.userName == "" || reflect.ValueOf(o.userName).IsZero() {
		return ErrUsernameNil
	}

	return nil
}

// Sign for generate signature request based on `username+apiKey+additional`
// with algorithm md5 hash
func (o *Option) Sign(c string) string {
	val := fmt.Sprintf("%s%s%s", o.userName, o.apiKey, c)
	data := []byte(val)

	hash := md5.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}
