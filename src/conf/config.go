package conf

import (
	"encoding/json"

	"fas/src/utils"
)

//配置对象
type Config struct {
	Title string `json:"title"`//系统标题
	Version string `json:"version"`//系统版本
	Port string `json:"port"`//服务端口
	Debug bool `json:"debug"`//是否为Debug模式
	LogPath string `json:"logPath"`//日志文件
	LogLevel string `json:"logLevel"`//日志级别

	Db DbConfig `json:"db"`//数据库配置
}

//数据库配置
type DbConfig struct {
	Type      string `json:"type"`//数据库类型
	Server    string `json:"server"`//服务器IP
	Port      int `json:"port"`//端口号
	Name      string `json:"name"`//数据库名称
	UserName  string `json:"user"`//数据库用户名
	Password  string `json:"pwd"`//数据库密码

	MaxIdleConns int `json:"maxIdle"`//最大空闲连接数
	MaxOpenConns int `json:"maxOpen"`//最大打开连接
}



//初始化配置文件
func (c *Config) InitConfig(fileName string)(bool, error){
	data, err := utils.LocalPathData(fileName)
	if err != nil {//读取配置文件失败
		return false, err
	}
	//解析为JSON对象
	e := json.Unmarshal(data, c)
	if e != nil {
		return false, e
	}
	return true, nil
}


