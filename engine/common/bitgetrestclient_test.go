package common

import (
	"testing"

	internal "github.com/zodiac163/bitget_api/engine"
)

func TestBitgetRestClient_HttpExecuter(t *testing.T) {
	restClient := new(BitgetRestClient).Init()
	params := internal.NewParams()
	params["productType"] = "umcbl"
	resp, _ := restClient.DoGet("/api/mix/v1/account/accounts", params)

	println(resp)
}
