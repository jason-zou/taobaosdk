package taobaosdk

import (
	"strconv"
)

type TopError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
    SubCode string `json:"sub_code"`
    SubMsg string `json:"sub_msg"`
}

func (e TopError) Error() string {
    return "Code: " + strconv.Itoa(e.Code) + ", Msg: " + e.Msg + ", SubCode: " + e.SubCode + ", SubMsg: " + e.SubMsg
}

type TaobaoErrResponse struct {
	ErrResponse      *TopError `json:"error_response"`
}

