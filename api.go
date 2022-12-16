package fapi_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type FinanceApiClient struct {
	baseUrl string
	apiKey  string
	client  http.Client
}

func (f *FinanceApiClient) GetTicker(symbol, interval, length string) (*TickerResponse, error) {
	path := fmt.Sprintf(`/v8/finance/chart/%s`, symbol)
	query := make(map[string]string)
	query["interval"] = interval
	query["range"] = length

	request, err := f.buildRequest(http.MethodGet, path, query)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := TickerResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (f *FinanceApiClient) GetQuote(symbol ...string) (*QuoteResponse, error) {
	path := "/v6/finance/quote"
	query := make(map[string]string)
	aggSym := symbol[0]

	if len(symbol) > 1 {
		aggSym = strings.Join(symbol, ",")
	}

	query["symbols"] = aggSym

	request, err := f.buildRequest(http.MethodGet, path, query)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := QuoteResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (f *FinanceApiClient) buildRequest(method, path string, query map[string]string) (*http.Request, error) {
	u := url.URL{
		Scheme:   "https",
		Host:     f.baseUrl,
		Path:     path,
		RawQuery: mapToQuery(query),
	}

	request, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add(X_API_KEY, f.apiKey)

	return request, err
}
