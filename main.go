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
	config := &conf.Config{}
	ok, err := config.InitConfig(confName)
	if !ok {
		fmt.Errorf("加载配置文件[%s]失败:%v", confName, err)
		return
	}
	fmt.Println(config.Title, ":", config.Version)
	//初始化日志
	log.InitLogger(config)
	//执行入口
	engine.Run(config)
}