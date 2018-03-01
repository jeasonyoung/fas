package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fas/src/conf"
	"fas/src/log"
	"fmt"
)

//数据库操作对象
var SqlDb *sql.DB

//初始化
func init(){
	log.GetLogInstance().Info("初始化连接数据库...")
	//
	dbConf := conf.GetConfigInstance().Db
	//检查数据库配置
	if dbConf == nil {
		log.GetLogInstance().Error("database config is null")
		return
	}
	//数据库类型
	if len(dbConf.Type) <= 0 {
		dbConf.Type = "mysql"
	}
	//数据库地址
	if len(dbConf.Server) <= 0 {
		dbConf.Server = "127.0.0.1"
	}
	//数据库端口
	if dbConf.Port <= 0 {
		dbConf.Port = 3306
	}
	//
	initDatabaseFromConf(dbConf)
	//
	log.GetLogInstance().Info("完成数据库连接")
}

//从配置加载数据
func initDatabaseFromConf(dbCof *conf.DbConfig){
	//拼接连接字符串
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%s?parseTime=true",
		dbCof.UserName, dbCof.Password, dbCof.Server, dbCof.Port, dbCof.Name)
	log.GetLogInstance().Info(fmt.Sprintf("数据库连接字符串:%v", conn))
	//
	var err error
	SqlDb, err = sql.Open(dbCof.Type, conn)
	//关闭数据库
	defer SqlDb.Close()
	//错误处理
	if err != nil {
		log.GetLogInstance().Error(err.Error())
		return
	}
	//
	maxIdle := 20
	maxOpen := 20
	if dbCof.MaxIdleConns > 0 {
		maxIdle = dbCof.MaxIdleConns
	}
	if dbCof.MaxOpenConns > 0 {
		maxOpen = dbCof.MaxOpenConns
	}
	SqlDb.SetMaxIdleConns(maxIdle)
	SqlDb.SetMaxOpenConns(maxOpen)
	//
	err = SqlDb.Ping()
	if err != nil {
		log.GetLogInstance().Error(err.Error())
	}
}