package rest

import (
    "strconv"
    "strings"
    "github.com/jason-zou/taobaosdk"
)

type ItemCatsGetRequest struct {
    taobaosdk.TaobaoRequest
}


func (t *ItemCatsGetRequest) SetFields(value string) {
    t.SetValue("fields", value)
}


func (t *ItemCatsGetRequest) SetParentCid(value int) {
    t.SetValue("parent_cid", strconv.Itoa(value))
}

func (t *ItemCatsGetRequest) SetCids(values []int) {
    vs := make([]string, 0)
    for _, v := range values {
        vs = append(vs, strconv.Itoa(v))
    }
    t.SetValue("cids", strings.Join(vs, ","))
}

func (t *ItemCatsGetRequest) GetResponse() (*ItemCatsGetResponse, error, *taobaosdk.TopError) {
    var resp ItemCatsGetResponseResult
    _, progErr, topErr := t.TaobaoRequest.GetResponse("taobao.itemcats.get", &resp, "")
    if progErr != nil {
        return nil, progErr, topErr
    }
    return resp.Response, progErr, topErr
}

