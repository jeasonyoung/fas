package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"

	"fas/src/engine"
	"fas/src/log"
	"errors"
)

//注册用户
type RegisterBody struct {
	Account string `json:"account"`//账号
	Password string `json:"password"`//密码
	Mobile string `json:"mobile"`//手机号码
	ValidCode string `json:"code"`//验证码
}

//解析报文体
func (body *RegisterBody) ParseBody(context *gin.Context) (bool, error){
	dataJson := context.MustGet(engine.ReqBodyJsonString).(string)
	log.Logger.Debug("ParseBody", zap.String(engine.ReqBodyJsonString, dataJson))
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

