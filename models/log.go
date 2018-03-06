package models

import "github.com/astaxie/beego/orm"

//终端日志
type ClientLog struct {
	Id string `orm:"column(id);pk"` //日志ID
	Mac string `orm:"column(mac)"`//设备标识
	UserId string `orm:"column(userId)"`//用户ID

	Type uint8 `orm:"column(type)"`//类型(0:normal,1:warn, 2:error)
	Path string `orm:"column(path)"`//日志文件路径
}

//表名称
func (c *ClientLog) TableName() string {
	return "tbl_fas_client_logs"
}

//初始化
func init(){
	orm.RegisterModel(new(ClientLog))
}