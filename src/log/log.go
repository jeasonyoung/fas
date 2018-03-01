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
	count := len(data)
	var fields = make([]zapcore.Field, count)
	//
	if count > 0 {
		for i := 0; i < count; i++ {
			d := data[i]
			f := zap.Any(d.Key, d.Value)
			fields = append(fields, f)
		}
	}
	//
	switch t {
		case logTypeDebug: {//debug
			l.logger.Debug(msg, fields...)
		}
		case logTypeInfo: {//info
			l.logger.Info(msg, fields...)
		}
		case logTypeWarn: {//warn
			l.logger.Warn(msg, fields...)
		}
		case logTypeFatal: {//fatal
			l.logger.Fatal(msg, fields...)
		}
		case logTypeError: {//error
			l.logger.Error(msg, fields...)
		}
	}
}