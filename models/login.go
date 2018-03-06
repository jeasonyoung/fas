package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//用户登录流水
type UserLogin struct {
	Id string `orm:"column(id);pk"`//登录ID
	UserId string `orm:"column(userId)"`//用户ID
	ChannelId string `orm:"column(channelId)"`//渠道ID
	Method uint8 `orm:"column(method)"`//登录方式(0:本地登录,1:微信,2:支付宝)

	Token string `orm:"column(token)"`//登录令牌
	IPAddr string `orm:"column(ipAddr)"`//登录IP地址
	Mac string `orm:"column(mac)"`//设备标识

	ExpiredTime time.Time `orm:"column(expiredTime)"`//过期时间戳
	Status uint8 `orm:"column(status)"`//状态(1:有效,0:无效)
}

//表名
func (u *UserLogin) TableName() string {
	return "tbl_fas_user_logins"
}

//用户登录流水历史
type UserLoginHistory struct {
	Id string `orm:"column(id);pk"`//登录ID
	UserId string `orm:"column(userId)"`//用户ID
	ChannelId string `orm:"column(channelId)"`//渠道ID
	Method string `orm:"column(method)"`//登录方式(0:本地登录,1:微信,2:支付宝)

	Token string `orm:"column(token)"`//登录令牌
	IPAddr string `orm:"column(ipAddr)"`//登录IP地址
	Mac string `orm:"column(mac)"`//设备标识
}

//表名
func (uh *UserLoginHistory) TableName() string {
	return "tbl_fas_user_login_histories"
}


//注册表
func init(){
	orm.RegisterModel(new(UserLogin), new(UserLoginHistory))
}