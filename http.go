package izanami_client

import (
	"io/ioutil"
	"net/http"
	"time"

	"fmt"
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

func (c *client) buildURL(path string, httpParams map[string]string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.hostname, path)
	req, errRequest := http.NewRequest(http.MethodGet, url, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	// Add query params
	q := req.URL.Query()
	for k, v := range httpParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set(headerClientID, c.clientID)
	req.Header.Set(headerClientSecret, c.clientSecret)

	return req, nil
}

func (c *client) get(path string, httpParams map[string]string) ([]byte, error) {
	req, errReq := c.buildURL(path, httpParams)
	if errReq != nil {
		return nil, errReq
	}

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
