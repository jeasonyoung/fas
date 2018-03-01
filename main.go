package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"fas/src/conf"
	"fas/src/log"
	"fas/src/engine"
)

//主程序入口
func main(){
	//
	fmt.Println("Gin Web System:", gin.Version)
	const confName = "conf.json"
	fmt.Println("load Config...", confName)
	//
	config := conf.GetConfigInstance()
	//初始化加载配置
	ok, err := config.InitConfig(confName)
	if !ok {
		fmt.Errorf("加载配置文件[%s]失败:%v", confName, err)
		return
	}
	//初始化日志
	log.GetLogInstance().InitLogger(config)
	//记录日志
	log.GetLogInstance().Info("welcome", log.Data("title", config.Title), log.Data("ver", config.Version))
	//执行入口
	engine.Run(config)
}