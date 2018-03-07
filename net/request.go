package net

import (
	"errors"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

//请求报文头
type ReqHead struct {
	Ver uint8 `json:"version"`//版本号
	Channel uint8 `json:"channel"`//渠道号
	Mac string `json:"mac"`//设备标识
	Token string `json:"token"`//令牌
	Time uint16 `json:"time"`//时间戳
	Sign string `json:"sign"`//签名戳
}

//请求报文
type Request struct {
	Head *ReqHead `json:"head"`//请求报文头
	Body interface{} `json:"body"`//请求报文体
}

//解析请求报文数据
func NewParseRequestBody(data []byte) (*Request, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	req := &Request{}
	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//解析报文体对象
func (req *Request) ToBodyTarget(target interface{}) error {
	if req.Body != nil && target != nil {
		//将body对象json
		data, err := json.Marshal(req.Body)
		if err != nil {
			logs.Debug("ToTargetBody[%v]-exp:%v", req.Body, err.Error())
			return err
		}
		//json解析
		if len(data) > 0 {
			return json.Unmarshal(data, target)
		}
	}
	return nil
}

//校验报文
func (req *Request) Verify() bool {
	//body := req.Body.(map[string]interface{})
	//

	return true
}

