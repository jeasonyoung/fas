package database

import (
	"errors"

	"github.com/satori/go.uuid"

	"go.uber.org/zap"

	"fas/src/log"
)

//用户
type User struct {
	Id string//用户ID
	Account string//用户账号
	Password string//密码

	NickName string//昵称
	IconUrl string//头像URL

	Mobile string//手机号码
	Email string //邮箱地址

	Status int//状态(1:启用,0:停用)
}

//ID是否存在
func (u *User) hasById(id string) bool {
	log.Logger.Debug("hasById", zap.String("id", id))
	if len(id) == 0 {
		return false
	}
	if SqlDb == nil {
		log.Logger.Debug("sql db is null")
		return false
	}
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_users where id=?", id).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false
	}
	return result
}

//账号是否存在
func (u *User) HasByAccount(account string) bool {
	log.Logger.Debug("HasByAccount", zap.String("account", account))
	if len(account) == 0 {
		return false
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false
	}
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_users where account=?", account).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false
	}
	return result
}

//手机号码是否存在
func (u *User) HasByMobile(mobile string) bool {
	log.Logger.Debug("HasByMobile", zap.String("mobile", mobile))
	if len(mobile) == 0 {
		return false
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false
	}
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_users where mobile=?", mobile).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false
	}
	return result
}

//email是否存在
func (u *User) HasByEmail(email string) bool {
	log.Logger.Debug("HasByEmail", zap.String("email", email))
	if len(email) == 0 {
		return false
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false
	}
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_users where email=?", email).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false
	}
	return result
}

//根据账号加载数据
func (u *User) LoadByAccount(account string)(bool,error){
	log.Logger.Debug("loadByAccount", zap.String("account", account))
	if len(account) == 0 {
		return false, errors.New("account is empty")
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//加载数据
	err := SqlDb.QueryRow("select id,account,password,nickName,iconUrl,mobile,email,status from tbl_fas_users where account=?", account).Scan(
		u.Id, u.Account, u.Password, u.NickName, u.IconUrl, u.Mobile, u.Email, u.Status)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	//
	return true,nil
}

//根据手机号码加载数据
func (u *User) LoadByMobile(mobile string)(bool, error){
	log.Logger.Debug("loadByMobile", zap.String("mobile", mobile))
	if len(mobile) == 0 {
		return false, errors.New("mobile is empty")
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//加载数据
	err := SqlDb.QueryRow("select id,account,password,nickName,iconUrl,mobile,email,status from tbl_fas_users where mobile=?", mobile).Scan(
		u.Id, u.Account, u.Password, u.NickName, u.IconUrl, u.Mobile, u.Email, u.Status)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	//
	return true,nil
}

//根据email加载数据
func (u *User) LoadByEmail(email string)(bool, error){
	log.Logger.Debug("loadByEmail", zap.String("email", email))
	if len(email) == 0 {
		return false, errors.New("email is empty")
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//加载数据
	err := SqlDb.QueryRow("select id,account,password,nickName,iconUrl,mobile,email,status from tbl_fas_users where email=?", email).Scan(
		u.Id, u.Account, u.Password, u.NickName, u.IconUrl, u.Mobile, u.Email, u.Status)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	//
	return true,nil
}

//根据类型和令牌加载数据
func (u *User) LoadByToken(token string,t int8)(bool, error){
	log.Logger.Debug("loadByToken", zap.String("token", token), zap.Int8("type", t))
	if len(token) == 0 {
		log.Logger.Fatal("loadByToken", zap.String("token", token))
		return false, nil
	}
	//
	err := SqlDb.QueryRow("select id,account,password,nickName,iconUrl,mobile,email,status from tbl_fas_users where id in (select userId from tbl_fas_user_oauths where type=? and authCode=?) limit 0,1", t, token).Scan(
		u.Id, u.Account, u.Password, u.NickName, u.IconUrl, u.Mobile, u.Email, u.Status)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, nil
	}
	return true,nil
}

//更新数据
func (u *User) SaveOrUpdate() (bool, error){
	log.Logger.Debug("SaveOrUpdate")
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(u.Id) == 0 {
		u.Id = uuid.NewV4().String()
	}
	//检查是否
	ret := u.hasById(u.Id)
	if !ret {//新增数据
		_, err := SqlDb.Exec("insert into tbl_fas_users(id,account,password,nickName,iconUrl,mobile,email,status) values(?,?,?,?,?,?,?,?)",
			u.Id, u.Account, u.Password, u.NickName, u.IconUrl, u.Mobile, u.Email, u.Status)
		if err != nil {
			log.Logger.Fatal(err.Error())
			return false, err
		}
		return true, nil
	}
	//更新数据
	_, err := SqlDb.Exec("update tbl_fas_users set account=?,password=?,nickName=?,iconUrl=?,mobile=?,email=?,status=? where id=?",
		u.Account, u.Password, u.NickName, u.IconUrl, u.Mobile, u.Email, u.Status, u.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}
