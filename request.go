package mobilepulsa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// APIRequest is abstract of HTTP Client that will make API call into iak.id backend
type APIRequest interface {
	Call(ctx context.Context, httpMethod, url string, header http.Header, body interface{}, result interface{}) error
}

// APIRequestImplementation is the default implementation of APIRequest
type APIRequestImplementation struct {
	HTTPClient *http.Client
}

func (a *APIRequestImplementation) Call(
	ctx context.Context,
	httpMethod,
	url string,
	header http.Header,
	body,
	result interface{},
) error {
	var (
		err     error
		reqBody []byte
	)
	
	hasBody := body != nil || (reflect.ValueOf(body).Kind() != reflect.Ptr && !reflect.ValueOf(body).IsNil())
	if hasBody {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	
	if header != nil {
		req.Header = header
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("mobilepulsa-go/%s", version))
	
	return a.doRequest(req, result)
}

func (a *APIRequestImplementation) doRequest(req *http.Request, result interface{}) error {
	res, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return ErrorHttp(resBody)
	}
	
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return err
	}
	
	return nil
}
