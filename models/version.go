package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//版本
type Version struct {
	Id string `orm:"column(id);pk"`//版本ID
	Name string `orm:"column(name)"`//版本名称
	Ver uint8 `orm:"column(version)"`//版本

	CheckCode string `orm:"column(checkCode)"`//校验码
	Status uint8 `orm:"column(status)"`//状态(1:有效,0:无效)

	StartTime time.Time `orm:"column(startTime)"`//生效时间
	Url string `orm:"column(url)"`//下载地址

	Desc string `orm:"column(description)"`//描述
	ChannelId string `orm:"column(channelId)"`//所属渠道ID
}

//表名
func (v *Version) TableName() string {
	return "tbl_fas_sys_versions"
}

//注册表
func init() {
	orm.RegisterModel(new(Version))
}
