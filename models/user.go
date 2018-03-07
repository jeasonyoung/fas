package models

import (
	"time"
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

const (
	StatusEnable = 0//启用
	StatusDisable = 1//停用
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
	//检查用户账号是否存在
	o := orm.NewOrm()
	//检查账号
	data := User{ Account:u.Account }
	err := o.Read(&data,"account")
	if err != nil {
		return errors.New("账号已存在")
	}
	//检查手机号码
	data = User{ Mobile:u.Mobile }
	err = o.Read(&data, "mobile")
	if err != nil {
		return errors.New("手机号码已注册")
	}
	//
	u.Id = uuid.NewV1().String()//用户ID
	u.Status = StatusEnable //状态
	//保存数据
	_, err = o.Insert(u)
	//
	return err
}

//表名
func (u *User) TableName() string {
	return "tbl_fas_users"
}


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
	orm.RegisterModel(new(User),new(UserLogin), new(UserLoginHistory))
}