package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

	_ "fas/routers"
)

//初始化函数，会在main之前被执行
func init() {
	//配置日志
	logs.SetLogger("console")
	//数据库配置
	dbAlias := beego.AppConfig.String("dbAlias")//数据库别名
	dbURL := beego.AppConfig.String("dbURL")
	//
	logs.Info("database[%v]:%v", dbAlias, dbURL)
	if len(dbAlias) > 0 && len(dbURL) > 0 {
		//数据库类型
		dbType := "mysql"
		//注册数据库类型
		orm.RegisterDriver(dbType, orm.DRMySQL)
		//注册数据库连接及别名
		orm.RegisterDataBase(dbAlias, dbType, dbURL)
		//启动Debug
		orm.Debug = beego.BConfig.RunMode == "dev"
	}
}

func main() {
	//api doc
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//添加过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		//cors
		ctx.Output.Header("Access-Control-Allow-Origin","*")
	})
	//执行入口
	beego.Run()
}
