package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//账本
type Account struct {
	Id string `orm:"column(id);pk"`//账本ID
	Code uint8 `orm:"column(code)"`//账本代码(排序)

	Name string `orm:"column(name)"`//账本名称
	Abbr string `orm:"column(abbr)"`//账本简称

	Type uint8 `orm:"column(type)"`//账本类型(0:私密账本,1:只读共享,2:公共账本)
	Status uint8 `orm:"column(status)"`//状态(0:封账,1:启用,2:删除)

	CreateUserId string `orm:"column(createUserId)"`//创建用户
}

//账本-表名
func (a *Account) TableName() string {
	return "tbl_fas_accounts"
}


//账本关联用户
type AccountUser struct {
	Id string `orm:"column(id);pk"`//关联ID
	AccountId string `orm:"column(accountId)"`//账本ID

	UserId string `orm:"column(userId)"`//用户ID
	Role uint8 `orm:"column(role)"`//共享角色(0:所有者,1:参与者)
}

//账本关联用户-表名
func (au *AccountUser) TableName() string {
	return "tbl_fas_account_users"
}

//账本明细
type AccountItem struct {
	Id string `orm:"column(id);pk"`//明细ID
	AccountId string `orm:"column(accountId)"`//所属账本ID
	Code uint8 `orm:"column(code)"`//账单序号

	UserId string `orm:"column(userId)"`//所属用户ID

	Title string `orm:"column(title)"`//名目
	Money float32 `orm:"column(money)"`//金额
	Time time.Time `orm:"column(time)"`//时间
}

//账本明细-表名
func (ai *AccountItem) TableName() string {
	return "tbl_fas_account_items"
}

//注册表
func init(){
	//
	orm.RegisterModel(new(Account), new(AccountUser), new(AccountItem))
}