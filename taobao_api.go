package taobaosdk

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func sign(secret string, params url.Values) string {
	joinedParams := ""
	keys := make([]string, 0)
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i := range keys {
		joinedParams += keys[i] + params[keys[i]][0]
	}
	joinedParams = secret + joinedParams + secret
	h := md5.New()
	h.Write([]byte(joinedParams))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum([]byte(""))))
}

type AppInfo struct {
	AppKey string
	Secret string
}

type TaobaoRequest struct {
	appInfo AppInfo
	reqUrl  string
	values  url.Values
}

func (t *TaobaoRequest) SetAppInfo(appKey string, secret string) {
	t.appInfo = AppInfo{appKey, secret}
}

func (t *TaobaoRequest) SetReqUrl(reqUrl string) {
	t.reqUrl = reqUrl
}

func (t *TaobaoRequest) GetReqUrl() string {
	return t.reqUrl
}

func (t *TaobaoRequest) SetValue(key, value string) {
	if t.values == nil {
		t.values = url.Values{}
	}
	t.values.Set(key, value)
}

func (t *TaobaoRequest) GetValue(key string) (string, error) {
	if t.values == nil {
		return "", errors.New("values are not set")
	}
	return t.values.Get(key), nil
}

func (t *TaobaoRequest) GetValues() url.Values {
	return t.values
}

func (t *TaobaoRequest) DelValue(key string) {
	t.values.Del(key)
}

func (t *TaobaoRequest) GetResponse(methodName string, resp interface{}, session string) ([]byte, error, *TopError) {
	t.SetReqUrl("http://gw.api.taobao.com/router/rest")
	t.SetValue("format", "json")
	t.SetValue("v", "2.0")
	t.SetValue("app_key", t.appInfo.AppKey)
	t.SetValue("sign_method", "md5")
	t.SetValue("timestamp", fmt.Sprint(time.Now().UnixNano()/1e6))
	t.SetValue("method", methodName)
	t.DelValue("sign")
	t.SetValue("sign", sign(t.appInfo.Secret, t.values))
	if session != "" {
		t.SetValue("session", session)
	}
    var topErr *TopError
	fmt.Println(t.GetValues().Encode())
	response, progErr := http.PostForm(t.GetReqUrl(), t.GetValues())
	fmt.Println(progErr)
	data, progErr := ioutil.ReadAll(response.Body)
	if progErr != nil {
		return data, progErr, topErr
	}
	var errResp TaobaoErrResponse
	progErr = json.Unmarshal(data, &errResp)
	if progErr != nil {
		return data, progErr, topErr
	}
	topErr = errResp.ErrResponse
	if topErr != nil {
        fmt.Println(topErr.Error())
		return data, progErr, topErr
	}
	return data, json.Unmarshal(data, resp), topErr
}
