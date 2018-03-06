package models

import "github.com/astaxie/beego/orm"

//渠道
type Channel struct {
	Id string `orm:"column(id);pk"`//渠道ID
	Code uint8 `orm:"column(code)"`//渠道代码
	Name string `orm:"column(name)"`//渠道名称
	Status uint8 `orm:"column(status)"`//状态(1:正常,0:停用)
}

//表名
func (c *Channel) TableName() string {
	return "tbl_fas_sys_channels"
}

//注册表
func init(){
	orm.RegisterModel(new(Channel))
}