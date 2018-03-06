package net


//响应类型码
type RespCode uint8

//响应状态码
const (
	RespCodeSuccess RespCode = 0//响应成功
	RespCodeParseRequestFail RespCode = 100//解析请求报文失败
	RespCodeVerifyRequestFail RespCode = 110//校验请求报文失败
	RespCodeGetRequestHeadFail RespCode = 120//获取报文头失败
	RespCodeChannelError RespCode = 130//渠道号错误
	RespCodeParamIsEmpty RespCode = 200//请求参数为空
	RespCodeDataValidError RespCode = 300//数据验证错误
	RespCodeDataStoreError RespCode = 310//数据存储错误
	RespCodeAccountNotExist RespCode = 400//账号不存在
	RespCodePasswordError RespCode = 410//密码错误
	RespCodeSignInFail RespCode = 420//登录失败
)

//响应状态码文本
var respCodeText = map[RespCode]string {
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
	Code uint8 `json:"code"`//响应码
	Msg string `json:"msg"`//响应消息
}

//初始化数据
func newRespHead(code RespCode, msg string) *RespHead {
	return &RespHead{Code:uint8(code), Msg:msg}
}

//响应报文
type Response struct {
	Head *RespHead `json:"head"`//响应报文头
	Body interface{} `json:"body"`//响应报文体
}

//设置消息头
func (rs *Response) SetHeadMessage(msg string) {
	if len(msg) == 0 {
		return
	}
	//设置消息头
	if rs.Head != nil {
		rs.Head.Msg = msg
	}
}

//初始化响应消息
func NewResponse(code RespCode) *Response {
	return &Response{ Head:newRespHead(code, respCodeText[code])}
}

//初始化响应消息
func NewResponseSuccess() *Response{
	return NewResponse(RespCodeSuccess)
}