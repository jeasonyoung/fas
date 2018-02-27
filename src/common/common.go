package common

const (
	ReqHead = "reqHead"//请求报文头
	ReqBodyJsonString = "reqBodyJsonString"//请求报文体
)

//令牌接口
type ITokenConf interface {
	//获取当前过期时间戳
	GetCurrentExpiredUnix() int64
}

//日志配置
type ILogConf interface {
	//获取是否Debug
	GetDebug() bool
	//日志文件路径
	GetLogPath() string
	//日志级别
	GetLogLevel() string
}