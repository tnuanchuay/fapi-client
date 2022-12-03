package fapi_client

import (
	"fmt"
	"testing"
)

func TestClientShouldGetResponseFromGetQuoteApi(t *testing.T) {
	expect := "AIRPORTS OF THAILAND PUBLIC COM"
	client := New("yfapi.net", "")
	resp, err := client.GetQuote("AOT.BK")
	if err != nil {
		t.Error(err)
	}
	actual := resp.QuoteResponse.Result[0].ShortName
	if expect != actual {
		t.Error(fmt.Sprintf("expect %s, actual %s", expect, actual))
	}
}

func TestGetTickerShouldReturnCorrectValue(t *testing.T) {
	client := New("yfapi.net", "")
	resp, err := client.GetTicker("KBANK.BK", "1d", "1y")
	if err != nil {
		t.Error(err)
	}

	if !(len(resp.Chart.Result) > 0) {
		t.Errorf("expect %s, actual %d", "more than 0", len(resp.Chart.Result))
	}
}
