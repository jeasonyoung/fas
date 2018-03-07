package models

import "github.com/astaxie/beego/orm"

//第三方登录
type OAuth struct {
	Id string `orm:"column(id);pk"`//关联ID
	UserId string `orm:"column(userId)"`//用户ID
	Type uint8 `orm:"column(type)"`//类型(1:wechat,2:alipay)
	AuthCode string `orm:"column(authCode)"`//第三方授权码
}

//表名
func (o *OAuth) TableName() string {
	return "tbl_fas_user_oauths"
}

//初始化
func init() {
	orm.RegisterModel(new(OAuth))
}