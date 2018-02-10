package log

import (
	"encoding/json"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"fas/src/conf"
)

var Logger *zap.Logger

//初始化日志
func InitLogger(conf *conf.Config) {
	//日志地址"out.log"自定义
	lp := conf.LogPath
	//日志级别DEBUG,ERROR,INFO
	lv := conf.LogLevel
	//是否DEBUG
	isDebug := true
	if conf.Debug != true {
		isDebug = false
	}
	initLogger(lp, lv, isDebug)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}

//初始化日志配置
func initLogger(logPath string, logLevel string, debug bool){
	var js string
	if debug {
		js = fmt.Sprintf(`{
			"level":"%s",
			"encoding":"json",
			"outputPaths":["stdout"],
			"errorOutputPaths":["stdout"]
		}`, logLevel)
	} else {
		js = fmt.Sprintf(`{
			"level":"%s",
			"encoding":"json",
			"outputPaths":["%s"],
			"errorOutputPaths":["%s"]
		}`, logLevel, logPath, logPath)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal("Init logger error:", err)
	}
}
