package models

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fas/src/log"
)

//响应状态码
const (
	RespCodeSuccess = 0//响应成功
	RespCodeParseRequestFail = 100//解析请求报文失败
	RespCodeVerifyRequestFail = 110//校验请求报文失败
	RespCodeGetRequestHeadFail = 120//获取报文头失败
	RespCodeChannelError = 130//渠道号错误
	RespCodeParamIsEmpty = 200//请求参数为空
	RespCodeDataValidError = 300//数据验证错误
	RespCodeDataStoreError = 310//数据存储错误
	RespCodeAccountNotExist = 400//账号不存在
	RespCodePasswordError = 410//密码错误
	RespCodeSignInFail = 420//登录失败
)

//响应状态码文本
var respCodeText = map[int]string {
	RespCodeSuccess : "响应成功",
	RespCodeParseRequestFail : "解析请求报文失败",
	RespCodeVerifyRequestFail : "校验请求报文失败",
	RespCodeGetRequestHeadFail : "获取报文头失败",
	RespCodeChannelError : "渠道号错误",
	RespCodeParamIsEmpty : "请求参数为空",
	RespCodeDataValidError : "数据验证错误",
	RespCodeDataStoreError : "数据存储错误",
	RespCodeAccountNotExist : "账号不存在",
	RespCodePasswordError : "密码错误",
	RespCodeSignInFail : "登录失败",
}

//响应头
type RespHead struct {
	Code int `json:"code"` //响应码
	Message string `json:"msg"`//响应消息
}

//响应报文
type Response struct {
	Head *RespHead `json:"head"`//响应报文头
	Body interface{} `json:"body"`//响应报文体
}

//初始化响应报文头
func (resp *Response) InitHead(code int, msg string){
	if resp.Head == nil {
		resp.Head = &RespHead{}
	}
	resp.Head.Code = code//状态码
	resp.Head.Message = msg//响应数据
}

//获取状态码文本
func (resp *Response) GetRespCodeText() string{
	if resp.Head == nil {
		return ""
	}
	return respCodeText[resp.Head.Code]
}

//检查请求参数
func (resp *Response) CheckReqParam(context *gin.Context, param string, err string, callback func(param string)(bool,error)) bool {
	log.GetLogInstance().Debug("checkReqParam", log.Data("param", param), log.Data("err", err))
	//检查字段是否为空
	if len(param) == 0 {
		log.GetLogInstance().Fatal(err)
		resp.InitHead(RespCodeParamIsEmpty, err)
		resp.ResponseJson(context)
		return false
	}
	//检查回调是否可用
	if callback != nil {
		//回调处理
		if ok, err := callback(param); !ok {
			log.GetLogInstance().Fatal(err.Error())
			resp.InitHead(RespCodeDataValidError, err.Error())
			resp.ResponseJson(context)
			return false
		}
	}
	return true
}

//解析报文失败响应输出
func (resp *Response) ResponseParseRequestFail(c *gin.Context, err error) {
	log.GetLogInstance().Debug("responseParseRequestFail", log.Data("err", err.Error()))
	resp.InitHead(RespCodeParseRequestFail, err.Error())
	resp.ResponseJson(c)
}

//输出JSON格式响应
func (resp *Response) ResponseJson(c *gin.Context){
	log.GetLogInstance().Debug("responseWrite")
	if c == nil {
		log.GetLogInstance().Fatal("context is nil")
		return
	}
	c.JSON(http.StatusOK, resp)
}