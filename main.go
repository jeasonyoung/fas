package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"fas/src/engine"
)


//主程序入口
func main(){
	//
	fmt.Println("Gin Web System:", gin.Version)
	fmt.Println("Family Accounting System ", engine.AppVersion)
	fmt.Println("load Config...")
	//
	engine.Run()
}