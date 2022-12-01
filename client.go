package fapi_client

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	X_API_KEY = "x-api-key"
)

func New(baseUrl, apiKey string) *FinanceApiClient {
	return &FinanceApiClient{
		Baseurl: baseUrl,
		ApiKey:  apiKey,
		client:  http.Client{},
	}
}

func mapToQuery(query map[string]string) string {
	arr := []string{}
	for k, v := range query {
		arr = append(arr, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(arr, "&")
}
