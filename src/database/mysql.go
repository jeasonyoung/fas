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

//初始化数据库处理
func InitDatabase(cfg *conf.Config){
	//检查数据库配置
	if cfg == nil || cfg.Db == nil {
		log.Logger.Error("Database config is null")
		return
	}
	//数据库
	dbConf := cfg.Db
	//数据库类型
	dbType := "mysql"
	if len(dbConf.Type) <= 0 {
		dbConf.Type = dbType
	}
	//数据库地址
	srvAddr := "127.0.0.1"
	if len(dbConf.Server) <= 0 {
		dbConf.Server = srvAddr
	}
	//数据库端口
	if dbConf.Port <= 0 {
		dbConf.Port = 3306
	}
	//数据库名
	if len(dbConf.Name) == 0 {
		log.Logger.Error("数据库名为空!")
		return
	}
	//账号
	if len(dbConf.UserName) == 0 {
		log.Logger.Error("数据账号为空!")
		return
	}
	//密码
	if len(dbConf.Password) == 0 {
		log.Logger.Error("数据库密码为空!")
		return
	}
	//
	initDatabaseFromConf(dbConf)
	//
	log.Logger.Info("完成数据库连接")
}

//从配置加载数据
func initDatabaseFromConf(dbCof *conf.DbConfig){
	//拼接连接字符串
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%s?parseTime=true",
		dbCof.UserName, dbCof.Password, dbCof.Server, dbCof.Port, dbCof.Name)
	log.Logger.Info(fmt.Sprintf("数据库连接字符串:%v", conn))
	//
	var err error
	SqlDb, err = sql.Open(dbCof.Type, conn)
	//关闭数据库
	defer SqlDb.Close()
	//错误处理
	if err != nil {
		log.Logger.Fatal(err.Error())
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
		log.Logger.Fatal(err.Error())
	}
}