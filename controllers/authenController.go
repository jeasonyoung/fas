package controllers

import (
	"fmt"
	"errors"

	"github.com/astaxie/beego/logs"

	"fas/net"
	"fas/models"
)

//认证控制器
type AuthenController struct {
	baseController
}

//POST 请求
func (ac *AuthenController) Post(){
	req := ac.GetReqData()
	logs.Debug("Post-req:%v", req)
	//初始化请求报文体
	reqBody := &ReqAuthenBody{}
	//解析请求报文
	err := req.ToBodyTarget(reqBody)
	if err != nil {
		ac.ResponseJsonWithParseRequestBodyFail(err)
		return
	}
	//校验请求数据
	err = reqBody.Valid()
	if err != nil {
		ac.ResponseJsonWithDataValidError(err)
		return
	}
	//设置渠道
	reqBody.channelId = fmt.Sprintf("%v", req.Head.Channel)
	//设置IP地址
	reqBody.ipAddr = ac.Ctx.Request.RemoteAddr
	//数据保存处理
	var respBody *RespAuthenBody
	respBody,err = reqBody.Save()
	if err != nil {
		ac.ResponseJsonWithDataStoreError(err)
		return
	}
	//响应处理
	ac.ResponseJson(net.NewResponseSuccessWithBody(respBody))
}

//用户验证-请求报文体
type ReqAuthenBody struct {
	Account   string `json:"account"`//账号
	Password  string `json:"password"`//密码
	Mac       string `json:"mac"`//设备标识
	channelId string //渠道ID
	ipAddr    string //IP地址
}

//校验数据
func (body *ReqAuthenBody) Valid() error {
	//检查账号
	if len(body.Account) == 0 {
		return errors.New("账号为空")
	}
	//检查密码
	if len(body.Password) == 0 {
		return errors.New("密码为空")
	}
	return nil
}

//保存数据
func (body *ReqAuthenBody) Save() (*RespAuthenBody, error) {
	u := &models.User{}
	//登录数据处理
	token, err := u.Sign(body.Account, body.Password, body.channelId, body.ipAddr, body.Mac, models.SignMethodWithLocal)
	if err != nil {
		return nil, err
	}
	//登录成功数据处理
	return &RespAuthenBody{
		Token:token,//登录令牌
		Info:RespUserInfo{
			NickName:u.NickName,//用户昵称
			IconURL:u.IconURL,//头像URL
		},
	},nil
}

//用户验证-响应报文体
type RespAuthenBody struct {
	Token string `json:"token"`//登录令牌
	Info RespUserInfo `json:"info"`//登录用户信息
}

//响应用户信息
type RespUserInfo struct {
	NickName string `json:"nickName"`//用户昵称
	IconURL  string `json:"iconUrl"`//头像URL
}