package log

import (
	"fmt"
	"log"
	"sync"
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"fas/src/common"
)
//
var lg *Log
var once sync.Once
//
//获取日志实例
func GetLogInstance() *Log {
	once.Do(func() {
		lg = &Log{}
	})
	return lg
}

//日志数据
type DataLog struct {
	Key string//键
	Value interface{}//值
}

//日志数据函数
func Data(key string, value interface{}) DataLog {
	return DataLog{ Key:key, Value: value }
}

//
type logType uint8
//
const (
	//debug
	logTypeDebug logType = iota
	//info
	logTypeInfo
	//warn
	logTypeWarn
	//fatal
	logTypeFatal
	//error
	logTypeError
)

//日志结构
type Log struct {
	logger *zap.Logger
}

//初始化日志
func (l *Log) InitLogger(cfg common.ILogConf) {
	//日志地址"out.log"自定义
	lp := cfg.GetLogPath()
	//日志级别DEBUG,ERROR,INFO
	lv := cfg.GetLogLevel()
	//是否DEBUG
	isDebug := true
	if cfg.GetDebug() != true {
		isDebug = false
	}
	l.initLogger(lp, lv, isDebug)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
}

//初始化日志配置
func (l *Log) initLogger(logPath string, logLevel string, debug bool){
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
	//
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal("Init logger error:", err)
		return
	}
	l.logger = logger
}

//debug
func (l *Log) Debug(msg string, data ... DataLog) {
	l.writeLog(logTypeDebug, msg, data)
}

//info
func (l *Log) Info(msg string, data ...DataLog) {
	l.writeLog(logTypeInfo, msg, data)
}

//warn
func (l *Log) Warn(msg string, data ...DataLog) {
	l.writeLog(logTypeWarn, msg, data)
}

//fatal
func (l *Log) Fatal(msg string, data ...DataLog) {
	l.writeLog(logTypeFatal, msg, data)
}

//error
func (l *Log) Error(msg string, data ...DataLog) {
	l.writeLog(logTypeError, msg, data)
}

//writeLog
func (l *Log) writeLog(t logType, msg string, data []DataLog){
	if l.logger == nil {
		return
	}
	//fmt.Printf("writeLog(logType:%v,msg:%v,data:%v)...\n", t, msg, data)
	//
	count := len(data)
	fields := make([]zapcore.Field, count)
	if count > 0 {
		for idx, d := range data {
			fmt.Printf("writeLog[%v]:%v\n", idx, d)
			fields[idx] = zap.Any(d.Key, d.Value)
		}
	}
	//fmt.Printf("writeLog-fields[%v]:%v\n", count, fields[:])
	//
	switch t {
		case logTypeDebug: {//debug
			if count > 0 {
				l.logger.Debug(msg, fields[:]...)
				return
			}
			l.logger.Debug(msg)
		}
		case logTypeInfo: {//info
			if count > 0 {
				l.logger.Info(msg, fields[:]...)
				return
			}
			l.logger.Info(msg)
		}
		case logTypeWarn: {//warn
			if count > 0 {
				l.logger.Warn(msg, fields[:]...)
				return
			}
			l.logger.Warn(msg)
		}
		case logTypeFatal: {//fatal
			if count > 0 {
				l.logger.Fatal(msg, fields[:]...)
				return
			}
			l.logger.Fatal(msg)
		}
		case logTypeError: {//error
			if count > 0 {
				l.logger.Error(msg, fields[:]...)
				return
			}
			l.logger.Error(msg)
		}
	}
}