package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fas/net"
)

//baseController 控制器基类
type baseController struct {
	beego.Controller
	req *net.Request
}

//Prepare 请求预处理
func (bc *baseController) Prepare() {
	logs.Info("Prepare...")
	bc.Data["xsrf_token"] = bc.XSRFToken()
	//解析请求报文
	req, err := net.NewParseRequestBody(bc.Ctx.Input.RequestBody)
	if err != nil {
		bc.req = nil
		logs.Warn(err.Error())
		return
	}
	//请求报文
	bc.req = req
}

//获取请求报文数据
func (bc *baseController) GetRequestData() *net.Request {
	return bc.req
}

//响应JSON数据
func (bc *baseController) ResponseJson(resp *net.Response) {
	if resp == nil {
		resp = net.NewResponseSuccess()
	}
	logs.Debug("Response-Json:%v", resp)
	bc.Data["json"] = resp
	bc.ServeJSON()
}

//响应成功处理
func (bc *baseController) ResponseJsonWithSuccess(){
	bc.ResponseJson(net.NewResponseSuccess())
}

//响应解析请求报文体失败
func (bc *baseController) ResponseJsonWithParseRequestBodyFail(err error) {
	resp := net.NewResponse(net.RespCodeParseRequestFail)
	if err != nil {
		resp.SetHeadMessage(err.Error())
	}
	bc.ResponseJson(resp)
}

//响应数据校验失败
func (bc *baseController) ResponseJsonWithDataValidError(err error) {
	resp := net.NewResponse(net.RespCodeDataValidError)
	if err != nil {
		resp.SetHeadMessage(err.Error())
	}
	bc.ResponseJson(resp)
}

//响应数据保存失败
func (bc *baseController) ResponseJsonWithDataStoreError(err error) {
	resp := net.NewResponse(net.RespCodeDataStoreError)
	if err != nil {
		resp.SetHeadMessage(err.Error())
	}
	bc.ResponseJson(resp)
}
