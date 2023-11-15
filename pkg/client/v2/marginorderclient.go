package v2

import (
	internal "github.com/zodiac163/bitget_api/engine"
	"github.com/zodiac163/bitget_api/engine/common"
)

type MarginOrderClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MarginOrderClient) Init() *MarginOrderClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

// normal order
func (p *MarginOrderClient) PlaceOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/margin/crossed/place-order", postBody)
	return resp, err
}

func (p *MarginOrderClient) BatchPlaceOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/margin/crossed/batch-place-order", postBody)
	return resp, err
}

func (p *MarginOrderClient) CancelOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/margin/crossed/cancel-order", postBody)
	return resp, err
}

func (p *MarginOrderClient) BatchCancelOrders(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/margin/crossed/batch-cancel-order", postBody)
	return resp, err
}

func (p *MarginOrderClient) OrdersHistory(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/margin/crossed/history-orders", params)
	return resp, err
}

func (p *MarginOrderClient) OrdersPending(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/margin/crossed/unfilled-orders", params)
	return resp, err
}

func (p *MarginOrderClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/margin/crossed/fills", params)
	return resp, err
}
