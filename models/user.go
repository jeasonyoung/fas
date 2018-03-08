package models

import (
	"time"
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/logs"
	"fas/utils"
	"strings"
)

type SignMethod uint8
type Status uint8

const (
	StatusDisable Status = iota//停用
	StatusEnable//启用
)

const (//登录方式
	SignMethodWithLocal SignMethod = iota//本地登录
	SignMethodWithWeChat //微信
	SignMethodWithAlipay //支付宝
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
	qs := o.QueryTable(u)
	//检查账号
	ok := qs.Filter("account", u.Account).Exist()
	if ok {
		return errors.New("账号已存在")
	}
	//检查手机号码
	ok = qs.Filter("mobile", u.Mobile).Exist()
	if ok {
		return errors.New("手机号码已注册")
	}
	//
	u.Id = uuid.NewV1().String()//用户ID
	u.Status = uint8(StatusEnable) //状态
	//保存数据
	_, err := o.Insert(u)
	//
	return err
}

//用户登录
func (u *User) Sign(account,passwd,channelId,ip,mac string, method SignMethod) (string, error){
	logs.Debug("Sign(account:%v,passwd:%v,channelId:%v,ip:%v,mac:%v,method:%v)...", account, passwd, channelId, ip, mac, method)
	o := orm.NewOrm()
	qs := o.QueryTable(u)
	//根据账号加载数据
	err := qs.Filter("account",account).One(u)
	if err != nil {
		return "", err
	}
	//检查账号是否有效
	if u.Status == uint8(StatusDisable) {
		return "", errors.New("账号已停用")
	}
	//校验密码
	encryptPwd := utils.MD5Sum(passwd)
	if !strings.EqualFold(u.Password, encryptPwd) {
		return "", errors.New("密码错误")
	}
	//登录流水
	login := &UserLogin{
		Id:uuid.NewV4().String(),//登录ID
		UserId:u.Id,//用户ID
		ChannelId:channelId,//渠道ID
		Method:uint8(method),//登录方式

		Token:uuid.NewV4().String(),//登录令牌
		IPAddr:ip,//登录IP地址
		Mac:mac,//设备标识

		Status:uint8(StatusEnable),//状态
	}
	//插入流水数据
	_, err = o.Insert(login)
	if err != nil {
		return "", err
	}
	//返回数据
	return login.Token, nil
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

	//ExpiredTime time.Time `orm:"column(expiredTime)"`//过期时间戳
	Status uint8 `orm:"column(status)"`//状态(1:有效,0:无效)
}

//表名
func (ul *UserLogin) TableName() string {
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