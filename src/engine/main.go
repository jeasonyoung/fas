package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


//当前版本
const AppVersion = "v1.0"

var (
	AppEngine *gin.Engine = nil
)

//初始化入口
func init() {
	AppEngine = gin.Default()
	//debug模式
	gin.SetMode(gin.DebugMode)
	//favicon.ico
	AppEngine.StaticFile("/favicon.ico", "./static/favicon.png")
	//static
	AppEngine.Static("/static", "./static")
}

//执行入口
func Run(){
	//
	AppEngine.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html;charset=utf-8")
		c.String(http.StatusOK, "<center><h1>%v %v</h1></center>", "FAS ", AppVersion)
	})

	//ping
	AppEngine.GET("/ping", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//启动http服务
	AppEngine.Run()
}