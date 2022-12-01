package fapi_client

import (
	"fmt"
	"testing"
)

func TestMapToQueryShouldReturnCorrectResultWhenEnter2Params(t *testing.T) {
	expect := "symbols=A,B,C&opt=A"
	input := map[string]string{
		"symbols": "A,B,C",
		"opt":     "A",
	}

	actual := mapToQuery(input)

	if expect != actual {
		t.Error(fmt.Sprintf("expect %s, actual %s", expect, actual))
	}
}

func TestMapToQueryShouldReturnCorrectResultWhenEnter1Params(t *testing.T) {
	expect := "symbols=A,B,C"

	input := map[string]string{
		"symbols": "A,B,C",
	}

	actual := mapToQuery(input)

	if expect != actual {
		t.Error(fmt.Sprintf("expect %s, actual %s", expect, actual))
	}
}
