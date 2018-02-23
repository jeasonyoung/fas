package models

import "github.com/gin-gonic/gin"

//解析请求报文接口
type RequestBodyParse interface{

	//解析报文体接口
	ParseBody(context *gin.Context) (bool,error)

}
