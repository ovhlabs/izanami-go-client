package izanami_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/facebookgo/httpcontrol"
)

const (
	httpParamPage     = "page"
	httpParamPageSize = "pageSize"

	headerClientID     = "Izanami-Client-Id"
	headerClientSecret = "Izanami-Client-Secret"
)

var (
	httpClient = &http.Client{
		Transport: &httpcontrol.Transport{
			RequestTimeout: time.Second * 30,
			MaxTries:       5,
		},
	}
)

// Metadata represents metadata parts of http response
type Metadata struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Count    int `json:"count"`
	NbPages  int `json:"nbPages"`
}

func (c *client) buildURL(path string, method string, httpParams map[string]string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.hostname, path)
	req, errRequest := http.NewRequest(method, url, body)
	if errRequest != nil {
		return nil, errRequest
	}

	if httpParams != nil {
		// Add query params
		q := req.URL.Query()
		for k, v := range httpParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headerClientID, c.clientID)
	req.Header.Set(headerClientSecret, c.clientSecret)

	return req, nil
}

func (c *client) get(path string, httpParams map[string]string) ([]byte, error) {
	req, errReq := c.buildURL(path, http.MethodGet, httpParams, nil)
	if errReq != nil {
		return nil, errReq
	}
	return do(req)
}

func (c *client) post(path string, body interface{}) ([]byte, error) {
	b, errJSON := json.Marshal(body)
	if errJSON != nil {
		return nil, errJSON
	}
	req, errReq := c.buildURL(path, http.MethodPost, nil, bytes.NewReader(b))
	if errReq != nil {
		return nil, errReq
	}
	return do(req)
}

func (c *client) put(path string, body interface{}) ([]byte, error) {
	b, errJSON := json.Marshal(body)
	if errJSON != nil {
		return nil, errJSON
	}
	req, errReq := c.buildURL(path, http.MethodPut, nil, bytes.NewReader(b))
	if errReq != nil {
		return nil, errReq
	}
	return do(req)
}

func (c *client) delete(path string) error {
	req, errReq := c.buildURL(path, http.MethodDelete, nil, nil)
	if errReq != nil {
		return errReq
	}
	_, errDo := do(req)
	return errDo
}

func do(req *http.Request) ([]byte, error) {
	res, errDo := httpClient.Do(req)
	if errDo != nil {
		return nil, errDo
	}
	defer res.Body.Close()
	body, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("[%d] %s", res.StatusCode, string(body))
	}
	return body, nil
}
