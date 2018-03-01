package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fas/src/conf"

	"fas/src/apis"
)

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
func Run(cfg *conf.Config){
	//版本v1
	v1 := AppEngine.Group("api/v1")
	{
		//使用报文解析中间件
		v1.Use(ParseRequestMiddleWare())
		{
			//通用API接口
			common := v1.Group("/common")
			{
				//初始化通用API
				commonApi := &apis.Common{}
				commonApi.InitToken(cfg)//初始化令牌配置数据
				//用户注册
				common.POST("/register", commonApi.Register)
				//用户登录
				common.POST("/signin", commonApi.SignIn)
			}
			//
		}
	}

	//
	AppEngine.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html;charset=utf-8")
		c.String(http.StatusOK, "<center><h1>%v %v</h1></center>", cfg.Title, cfg.Version)
	})

	//ping
	AppEngine.GET("/ping", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//启动http服务
	AppEngine.Run(cfg.Port)
}