package models

import (
	"errors"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"

	"fas/src/log"
	"fas/src/common"
)

//登录报文体
type SignInBody struct {
	Account string `json:"account"`//登录账号
	Password string `json:"password"`//登录密码
}

//解析登录报文体
func (body *SignInBody) ParseBody(context *gin.Context) (bool,error){
	dataJson := context.MustGet(common.ReqBodyJsonString).(string)
	log.Logger.Debug("parseBody", zap.String(common.ReqBodyJsonString, dataJson))
	if len(dataJson) == 0 {
		return false, errors.New("json is empty")
	}
	//JSON解析
	err := json.Unmarshal([]byte(dataJson), body)
	if err != nil {
		log.Logger.Error(err.Error())
		return false, err
	}
	return true, nil
}

//登陆结果响应报文体
type SignResultBody struct {
	Token string `json:"token"`//登录令牌
	Name string `json:"name"`//用户名称
	IconUrl string `json:"iconUrl"`//头像URL
}