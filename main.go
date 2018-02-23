package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"

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
	config := &conf.Config{}
	ok, err := config.InitConfig(confName)
	if !ok {
		fmt.Errorf("加载配置文件[%s]失败:%v", confName, err)
		return
	}
	//初始化日志
	log.InitLogger(config)
	//
	log.Logger.Info("welcome", zap.String("title", config.Title), zap.String("ver", config.Version))
	//执行入口
	engine.Run(config)
}