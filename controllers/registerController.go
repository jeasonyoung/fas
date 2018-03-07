package controllers

import (
	"errors"

	"github.com/astaxie/beego/logs"

	"fas/models"
	"fas/utils"
)

//注册控制器
type RegisterController struct {
	baseController
}

//POST 请求
func (rc *RegisterController) Post() {
	req := rc.GetReqData()
	logs.Debug("Post-req:%v", req)
	//初始化请求报文体
	reqBody := &ReqRegisterBody{}
	//解析请求报文
	err := req.ToBodyTarget(reqBody)
	if err != nil {
		rc.ResponseJsonWithParseRequestBodyFail(err)
		return
	}
	//请求报文处理
	rc.ReqBodyHandler(reqBody)
}

//注册用户-请求报文体
type ReqRegisterBody struct {
	Account   string `json:"account"`//账号
	Password  string `json:"password"`//密码
	Mobile    string `json:"mobile"`//手机号码
}

//校验数据
func (body *ReqRegisterBody) Valid() error {
	//检查账号
	if len(body.Account) == 0 {
		return errors.New("账号为空")
	}
	//检查密码
	if len(body.Password) == 0 {
		return errors.New("密码为空")
	}
	//检查手机号码
	if len(body.Mobile) == 0 {
		return errors.New("手机号码为空")
	}
	return nil
}

//保存数据
func (body *ReqRegisterBody) Save() error {
	//初始化用户
	user := &models.User{
		Account:body.Account,//账号
		Password:utils.MD5Sum(body.Password),//密码
		Mobile:body.Mobile,//手机号码
	}
	//保存注册用户
	return user.Register()
}


