package models

import (
	"github.com/astaxie/beego/orm"
)

//用户
type User struct {
	Id string `orm:"column(id);pk"`//用户ID
	Account string `orm:"column(account)"`//用户账号
	Password string `orm:"column(password)"`//用户密码

	NickName string `orm:"column(nickName)"`//用户昵称
	IconURL string `orm:"column(iconUrl)"`//头像URL

	Mobile string `orm:"column(mobile)"`//手机号码
	Email string `orm:"column(email)"`//邮件地址

	Status uint8 `orm:"column(status)"`//状态(1:启用,0:停用)
}

//用户注册
func (u *User) Register() error {


	return nil
}

//表名
func (u *User) TableName() string {
	return "tbl_fas_users"
}

//注册表
func init(){
	orm.RegisterModel(new(User))
}