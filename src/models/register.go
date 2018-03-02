package models

import (
	"errors"

	"encoding/json"

	"github.com/gin-gonic/gin"

	"fas/src/log"

	"fas/src/common"
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
	dataJson := context.MustGet(common.ReqBodyJsonString).(string)
	log.GetLogInstance().Debug("ParseBody", log.Data(common.ReqBodyJsonString, dataJson))
	if len(dataJson) == 0 {
		return false, errors.New("json is empty")
	}
	//JSON解析
	err := json.Unmarshal([]byte(dataJson), body)
	if err != nil {
		log.GetLogInstance().Error(err.Error())
		return false, err
	}
	return true, nil
}

