package engine

import (
	"github.com/gin-gonic/gin"

	"fas/src/log"
	"fas/src/models"
	"fas/src/common"
)

//请求报文解析中间件
func ParseRequestMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求报文数据
		req := &models.Request{}
		//解析数据
		ret, err := req.ParseRequest(context)
		if !ret || err != nil {
			log.GetLogInstance().Error("解析请求报文失败:", log.Data("err", err.Error()))
			//初始化响应输出报文
			resp := &models.Response{}
			resp.InitHead(models.RespCodeParseRequestFail, "解析请求报文失败:" + err.Error())
			resp.ResponseJson(context)
			//
			context.Abort()
			return
		}
		//
		log.GetLogInstance().Debug("request-json:", log.Data("result", ret), log.Data("data", req))
		//验证报文
		if !req.Verify() {
			log.GetLogInstance().Error("校验请求报文失败")
			//初始化响应输出报文
			resp := &models.Response{}
			resp.InitHead(models.RespCodeVerifyRequestFail, "校验请求报文失败")
			resp.ResponseJson(context)
			//
			context.Abort()
			return
		}
		//设置请求报文头
		context.Set(common.ReqHead, req.Head)
		//设置请求报文体
		context.Set(common.ReqBodyJsonString, req.BodyToJsonString())
		//
		context.Next()
	}
}