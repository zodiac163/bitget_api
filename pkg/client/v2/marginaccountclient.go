package v2

import (
	internal "github.com/zodiac163/bitget_api/engine"
	"github.com/zodiac163/bitget_api/engine/common"
)

type MarginAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MarginAccountClient) Init() *MarginAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

// func (p *MarginAccountClient) Info() (string, error) {
// 	params := internal.NewParams()
// 	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/info", params)
// 	return resp, err
// }

func (p *MarginAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/margin/crossed/account/assets", params)
	return resp, err
}

func (p *MarginAccountClient) Repay(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/margin/crossed/account/repay", postBody)
	return resp, err
}

func (p *MarginAccountClient) MaxTransferable(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/margin/crossed/account/max-transfer-out-amount", params)
	return resp, err
}

// func (p *MarginAccountClient) Bills(params map[string]string) (string, error) {
// 	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/bills", params)
// 	return resp, err
// }

func (p *MarginAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/transferRecords", params)
	return resp, err
}
