package models

import (
	"errors"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"fas/src/log"
	"fas/src/common"
)

//请求报文头
type ReqHead struct {
	Version uint16 `json:"version"`//版本号
	Channel uint8 `json:"channel"`//渠道代码
	Mac string `json:"mac"`//设备标识
	Token string `json:"token"`//令牌
	Time uint64 `json:"time"`//时间戳
	Sign string `json:"sign"`//签名戳
}

//解析获取请求报文头
func (head *ReqHead) ParseRequest(context *gin.Context)(bool, error){
	log.GetLogInstance().Debug("parseRequest")
	if context == nil {
		log.GetLogInstance().Error("context is null")
		return false, errors.New("context is null")
	}
	//获取消息头数据
	data := context.MustGet(common.ReqHead).(*ReqHead)
	if data == nil {
		log.GetLogInstance().Debug("中间件未设置消息头数据")
		return false, errors.New("中间件未设置消息头数据")
	}
	//重新赋值
	head.Version = data.Version//版本号
	head.Channel = data.Channel//渠道号
	head.Token = data.Token//令牌
	head.Time = data.Time//时间戳
	head.Sign = data.Sign//签名戳
	return true, nil
}

//请求报文
type Request struct {
	Head *ReqHead `json:"head"`//请求报文头
	Body interface{} `json:"body"`//请求报文体
}

//解析请求数据
func (req *Request) ParseRequest(context *gin.Context)(bool, error){
	log.GetLogInstance().Debug("parseRequest...")
	if context == nil {
		log.GetLogInstance().Error("context is null")
		return false, errors.New("context is null")
	}
	//解析报文为数据json
	err := context.BindJSON(req)
	if err != nil {
		log.GetLogInstance().Error(err.Error())
		return false, err
	}
	return true, nil
}

//将body解析为json字符串
func (req *Request) BodyToJsonString() string {
	log.GetLogInstance().Debug("BodyToJsonString", log.Data("body", req.Body))
	if req.Body != nil {
		data, err := json.Marshal(req.Body)
		if err != nil {
			log.GetLogInstance().Fatal(err.Error())
			return "{}"
		}
		return string(data)
	}
	return "{}"
}

//校验报文
func (req *Request) Verify() bool {
	body := req.Body.(map[string]interface{})
	log.GetLogInstance().Info("verify", log.Data("head", req.Head), log.Data("body", body))

	//t := reflect.TypeOf(req.Body)
	//log.Logger.Debug("verify-type", zap.Any("type", t))
	///TODO:
	return true
}