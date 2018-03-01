package apis

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"

	"fas/src/common"

	"fas/src/log"
	"fas/src/models"

	"fas/src/utils"

	"fas/src/dao"
)

//公用结构体
type Common struct {
	cToken common.ITokenConf
}

//初始化配置数据
func (c *Common) InitToken(token common.ITokenConf){
	if token == nil {
		return
	}
	c.cToken = token
}

//用户注册
func (c *Common) Register(context *gin.Context) {
	log.GetLogInstance().Debug("register")
	//初始化响应报文
	resp := &models.Response{}
	//初始化用户注册报文体
	body := &models.RegisterBody{}
	//解析报文体
	_, err := body.ParseBody(context)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		resp.ResponseParseRequestFail(context, err)
		return
	}
	//检查密码
	if !resp.CheckReqParam(context, body.Password, "密码为空", nil) {
		return
	}
	//初始化用户数据
	user := &dao.User{}
	//检查账号
	if !resp.CheckReqParam(context, body.Account, "账号为空", func(param string) (bool, error) {
		//检查账号是否已存在
		if user.HasByAccount(body.Account) {
			log.GetLogInstance().Fatal("account is exists")
			return false, errors.New("账号已存在")
		}
		return true, nil
	}) {
		return
	}
	//检查手机号码
	if !resp.CheckReqParam(context, body.Mobile, "手机号码为空", func(param string) (bool, error) {
		//检查手机号码是否已存在
		if user.HasByMobile(body.Mobile) {
			log.GetLogInstance().Fatal("mobile is exists")
			return false, errors.New("手机号码已注册")
		}
		return true,nil
	}){
		return
	}
	//检查验证码
	if !resp.CheckReqParam(context, body.ValidCode, "验证码为空", nil){
		return
	}
	//保存数据
	user.Account = body.Account//账号
	user.Password = utils.MD5Sum(body.Password)//密码
	user.Mobile = body.Mobile//手机号码
	user.Status = 1//启用
	//保存数据
	ret,err := user.SaveOrUpdate()
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		resp.InitHead(models.RespCodeDataStoreError, err.Error())
	}else {
		retCode := models.RespCodeSuccess
		info := "保存数据成功"
		if !ret {
			retCode = models.RespCodeDataStoreError
			info = "保存数据失败"
		}
		log.GetLogInstance().Info("register-result:", log.Data("retCode", retCode))
		resp.InitHead(int(retCode), info)
	}
	resp.ResponseJson(context)
}

//用户登录
func (c *Common) SignIn(context *gin.Context) {
	log.GetLogInstance().Debug("signIn...")
	//初始化响应报文
	resp := &models.Response{}
	//初始化用户登录报文体
	body := &models.SignInBody{}
	//解析请求报文体
	_, err := body.ParseBody(context)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		resp.ResponseParseRequestFail(context, err)
		return
	}
	//检查账号是否为空
	if !resp.CheckReqParam(context, body.Account, "账号为空", nil) {
		return
	}
	//检查密码是否为空
	if !resp.CheckReqParam(context, body.Password, "密码为空", nil) {
		return
	}
	//初始化用户数据
	user := &dao.User{}
	//通过账号加载用户数据
	_, err = user.LoadByAccount(body.Account)
	if err != nil {
		log.GetLogInstance().Fatal("账号不存在")
		resp.InitHead(models.RespCodeAccountNotExist, "账号不存在")
		resp.ResponseJson(context)
		return
	}
	//检查密码是否错误
	log.GetLogInstance().Debug("input passwd:", log.Data("account", body.Account), log.Data("password", body.Password))
	encryptPasswd := utils.MD5Sum(body.Password)
	//对比密码
	if strings.ToLower(encryptPasswd) == strings.ToLower(user.Password) {
		//获取消息头
		reqHead := &models.ReqHead{}
		_, err = reqHead.ParseRequest(context)
		if err != nil {
			resp.InitHead(models.RespCodeGetRequestHeadFail, err.Error())
			resp.ResponseJson(context)
			return
		}
		//加载渠道数据
		channel := &dao.Channel{}
		_, err = channel.LoadByCode(int(reqHead.Channel))
		if err != nil {
			resp.InitHead(models.RespCodeChannelError, err.Error())
			resp.ResponseJson(context)
			return
		}
		//生成登陆令牌
		token := utils.MD5Sum(uuid.NewV4().String())
		//
		userLogin := &dao.UserLogin{
			UserId: user.Id,//用户ID
			ChannelId:channel.Id,//渠道ID
			Method:0,//登录方式(0:本地登录,1:微信,2:支付宝)
			Token:token,//登录令牌
			IpAddr:context.Request.RemoteAddr,//请求IP地址
			Mac:reqHead.Mac,//设备标识
			ExpiredTime:c.cToken.GetCurrentExpiredUnix(),//过期时间戳
			Status:1,//状态(1:有效,0:无效)
		}
		//令牌保存
		_, err = userLogin.SaveOrUpdate()
		if err != nil {
			log.GetLogInstance().Fatal(err.Error())
			resp.InitHead(models.RespCodeSignInFail, err.Error())
			resp.ResponseJson(context)
			return
		}
		//登录成功
		resp.InitHead(models.RespCodeSuccess,"登录成功")
		resp.Body = &models.SignResultBody{
			Token:token,//登录令牌
			Name:user.NickName,//用户昵称
			IconUrl:user.IconUrl,//头像URL
		}
	} else {
		//密码错误
		resp.InitHead(models.RespCodePasswordError, "密码错误")
	}
	//响应输出
	resp.ResponseJson(context)
}
