package rest

import (
	"github.com/jason-zou/taobaosdk"
)

type ShopGetRequest struct {
	taobaosdk.TaobaoRequest
}

func (t *ShopGetRequest) SetFields(value string) {
	t.SetValue("fields", value)
}

func (t *ShopGetRequest) SetNick(value string) {
	t.SetValue("nick", value)
}

func (t *ShopGetRequest) GetResponse() (*ShopGetResponse, error, *taobaosdk.TopError) {
	var resp ShopGetResponseResult
	_, progErr, topErr := t.TaobaoRequest.GetResponse("taobao.shop.get", &resp, "")
	if progErr != nil {
		return nil, progErr, topErr
	}
	return resp.Response, progErr, topErr
}
