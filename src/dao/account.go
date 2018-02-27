package dao

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"github.com/satori/go.uuid"

	"fas/src/log"
)

//账本
type Account struct {
	Id string//账本ID
	Code uint32//账本代码
	Name string//账本名称
	Abbr string//账本简称
	Type uint8//账本类型(0:私密账本,1:只读共享,2:公共账本)
	Status uint8//状态(0:封装,1:启用,2:删除)
	CreateUserId string//创建用户ID
}

//账本关联用户
type AccountUser struct {
	Id string//关联ID
	AccountId string//账本ID
	UserId string//用户ID
	Role uint8//共享角色(0:所有者,1:参与者)
}

//账本明细
type AccountItem struct {
	Id string//明细ID
	AccountId string//所属账本ID
	Code uint64//账单序号

	UserId string//所属用户ID

	Title string//名目
	Money float32//金额
	Time time.Time//时间
}

//检查数据
func (a *Account) hasById(id string)(bool, error){
	log.Logger.Debug("hasById", zap.String("id", id))
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(id) == 0 {
		log.Logger.Debug("id is empty")
		return false, errors.New("id is empty")
	}
	//
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_accounts where id=?", id).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false,err
	}
	return result, nil
}

//根据账本ID
func (a *Account) LoadById(id string)(bool, error){
	log.Logger.Debug("LoadById", zap.String("id", id))
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(id) == 0 {
		log.Logger.Fatal("id is empty")
		return false, errors.New("id is empty")
	}
	//
	err := SqlDb.QueryRow("select id,code,name,abbr,type,status,createUserId from tbl_fas_accounts where id=?", id).Scan(
		a.Id, a.Code, a.Name, a.Abbr, a.Type, a.Status, a.CreateUserId)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//新增或保存数据
func (a *Account) SaveOrUpdate()(bool, error){
	log.Logger.Debug("saveOrUpdate")
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(a.Id) == 0 {
		a.Id = uuid.NewV4().String()
	}
	ret, err := a.hasById(a.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	if !ret {//新增
		_, err = SqlDb.Exec("insert into tbl_fas_accounts(id,code,name,abbr,type,status,createUserId) values(?,?,?,?,?,?,?)",
			a.Id, a.Code, a.Name, a.Abbr, a.Type, a.Status, a.CreateUserId)
		if err != nil {
			log.Logger.Fatal(err.Error())
			return false, err
		}
		return true, nil
	}
	//更新
	_, err = SqlDb.Exec("update tbl_fas_accounts set code=?,name=?,abbr=?,type=?,status=?,createUserId=? where id=?",
		a.Code, a.Name, a.Abbr, a.Type, a.Status, a.CreateUserId, a.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//根据账本ID删除数据
func (a *Account) RemveById(id string)(bool, error){
	log.Logger.Debug("RemveById", zap.String("id", id))
	if len(id) == 0 {
		return false, errors.New("id is empty")
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//检查账本明细数据
	var ret bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_account_items where accountId=?", id).Scan(&ret)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	if ret {//账本明细已存在，则将账本状态设置为删除
		rs, err := SqlDb.Exec("update tbl_fas_accounts set status=2 where id=?", id)
		if err != nil {
			log.Logger.Fatal(err.Error())
			return false, err
		}
		count, _ := rs.RowsAffected()
		return count > 0, nil
	}
	//删除账本关联用户表
	_, err = SqlDb.Exec("delete from tbl_fas_account_users where accountId=?", id)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	//删除账本数据
	rs, err := SqlDb.Exec("delete from tbl_fas_accounts where id=?", id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	count, _ := rs.RowsAffected()
	return count > 0, nil
}

//新增或更新账本关联用户
func (au *AccountUser) SaveOrUpdate()(bool, error)  {
	log.Logger.Debug("saveOrUpdate")
	if len(au.UserId) == 0 || len(au.AccountId) == 0 {
		log.Logger.Fatal("userId or accountId is empty")
		return false, errors.New("userId or accountId is empty")
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//检查是否存在
	var exists bool
	err := SqlDb.QueryRow("select count(0)>0 from tbl_fas_account_users where accountId=? and userId=?", au.AccountId, au.UserId).Scan(&exists)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	if exists {
		//删除已存在数据
		_, err = SqlDb.Exec("delete from tbl_fas_account_users where accountId=? and userId=?", au.AccountId, au.UserId)
		if err != nil {
			log.Logger.Fatal(err.Error())
		}
	}
	//新增数据
	if len(au.Id) == 0 {
		au.Id = uuid.NewV4().String()
	}
	_, err = SqlDb.Exec("insert into tbl_fas_account_users(id,accountId,userId,role) values(?,?,?,?)", au.Id, au.AccountId, au.UserId, au.Role)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//检查数据是否存在
func (ai *AccountItem) hasById(id string)(bool, error){
	log.Logger.Debug("hasById", zap.String("id", id))
	if len(id) == 0 {
		log.Logger.Fatal("id is empty")
		return false, errors.New("id is empty")
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	var exits bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_account_items where id=?", id).Scan(&exits)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return exits, nil
}

//获取账本明细的最大账单序号
func (ai *AccountItem) loadMaxCode(accountId string) uint64 {
	log.Logger.Debug("loadMaxCode", zap.String("accountId", accountId))
	if len(accountId) == 0 {
		log.Logger.Fatal("accountId is empty")
		return 0
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return 0
	}
	var code uint64
	err := SqlDb.QueryRow("select max(code) from tbl_fas_account_items where accountId=?", accountId).Scan(&code)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return 0
	}
	return code
}

//新增或更新账本明细
func (ai *AccountItem) SaveOrUpdate() (bool, error){
	log.Logger.Debug("saveOrUpdate")
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(ai.AccountId) == 0 {
		log.Logger.Fatal("accountId is empty")
		return false, errors.New("accountId is empty")
	}
	if len(ai.UserId) == 0 {
		log.Logger.Fatal("userId is empty")
		return false, errors.New("userId is empty")
	}
	add := false
	if len(ai.Id) == 0 {
		add = true
		ai.Id = uuid.NewV4().String()
	}
	if !add {
		//检查数据是否存在
		ret, _ := ai.hasById(ai.Id)
		if !ret {
			add = true
		}
	}
	var err error
	if add {//新增数据
		//加载最大排序号
		maxCode := ai.loadMaxCode(ai.AccountId)
		ai.Code = maxCode + 1
		//
		_, err = SqlDb.Exec("insert into tbl_fas_account_items(id,accountId,code,userId,title,money,time) values(?,?,?,?,?,?,?)",
			ai.Id, ai.AccountId, ai.Code, ai.UserId, ai.Title, ai.Money, ai.Time)
	}else {
		//更新数据
		_, err = SqlDb.Exec("update tbl_fas_account_items set accountId=?,userId=?,title=?,money=?,time=? where id=?",
			ai.AccountId, ai.UserId, ai.Title, ai.Money, ai.Time, ai.Id)
	}
	//结果数据处理
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true,nil
}