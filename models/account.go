package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type AccountType uint8//账本类型
type AccountStatus uint8//账本状态
type AccountRole uint8//账本角色

const (
	AccountTypeWithAll AccountType = iota//全部账本
	AccountTypeWithPrivate//私人账本
	AccountTypeWithRead//只读账本
	AccountTypeWithPublic//公共账本
)

const (
	AccountStatusWithDisable AccountStatus = iota//封账
	AccountStatusWithEnable//启用
	AccountStatusWithDelete//删除
)

const (
	AccountRoleWithOwner AccountRole = iota//所有者
	AccountRoleWithVisitor //参与者
)

//账本信息
type AccountInfo struct {
	Id string//账本ID
	Code uint8//账本代码
	Name string//账本名称
	Abbr string//账本简称
	Type uint8//账本类型
	Status uint8//账本状态
	Role uint8//账本角色
}


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

//查询账本数据
func (a *Account) QueryAccounts(userId string,tp,status uint8, index, rows uint)(uint, *[]AccountInfo){
	logs.Debug("QueryAccounts(userId:%v,type:%v,status:%v,index:%v,rows:%v)...", userId, tp, status, index, rows)
	//初始化查询
	o := orm.NewOrm()
	//查询总数据量
	var totals uint
	err := o.Raw("select count(0) from vw_fas_accounts where userId = ? and type = ? and status = ?", userId, tp, status).QueryRow(&totals)
	if err != nil {
		logs.Warn("QueryAccounts(userId:%v,type:%v,status:%v):%v", userId, tp, status,err.Error())
		return 0, nil
	}
	//
	start := (index - 1) * rows
	var items []AccountInfo
	_, err = o.Raw("select id,code,name,abbr,type,status,role from vw_fas_accounts where userId = ? and type = ? and status = ? limit ?,?",
		userId, tp, status, start, rows).QueryRows(&items)
	if err != nil {
		logs.Warn("QueryAccounts(userId:%v,type:%v,status:%v):%v", userId, tp, status,err.Error())
		return 0, nil
	}
	return totals, &items
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