package database

import (
	"go.uber.org/zap"

	"fas/src/log"
	"errors"
	"github.com/satori/go.uuid"
)


//第三方登录关联
type OAuth struct {
	Id string //关联ID
	UserId string//用户ID
	Type int8//类型(1.wechat,2:alipay)
	AuthCode string//授权码
}

//检查类型和代码是否存在
func hasByTypeAndCode(t int8, code string) (bool, error) {
	log.Logger.Debug("HasByTypeAndCode", zap.Int8("type", t), zap.String("code", code))
	if len(code) == 0 {
		log.Logger.Debug("code is empty")
		return false, errors.New("code is empty")
	}
	if SqlDb == nil {
		log.Logger.Debug("sql db is null")
		return false, errors.New("sql db is null")
	}
	//
	var result bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_user_oauths where type=? and authCode=?", t, code).Scan(&result)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return result, nil
}

//新增或更新数据
func (o * OAuth) SaveOrUpdate() (bool, error) {
	log.Logger.Debug("SaveOrUpdate")
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(o.Id) == 0 {
		o.Id = uuid.NewV4().String()
	}
	//检查是否有重复
	ret, err := hasByTypeAndCode(o.Type, o.AuthCode)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	if ret {//
		return false, errors.New("授权码已存在")
	}
	//新增
	_, err = SqlDb.Exec("insert into tbl_fas_user_oauths(id,userId,type,authCode) values(?,?,?,?)", o.Id, o.UserId, o.Type, o.AuthCode)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//根据类型和授权码
func (o *OAuth) RemoveByTypeAndCode(t int8, code string)(bool, error){
	log.Logger.Debug("RemoveByTypeAndCode", zap.Int8("type", t), zap.String("authCode", code))
	if len(code) == 0 {
		log.Logger.Fatal("authCode is empty")
		return false, errors.New("authCode is empty")
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//删除数据
	_, err := SqlDb.Exec("delete from tbl_fas_user_oauths where type=? and authCode=?", t, code)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//根据用户删除数据
func (o *OAuth) RemoveByUser(userId string)(bool, error){
	log.Logger.Debug("RemoveByUser", zap.String("userId", userId))
	if len(userId) == 0 {
		log.Logger.Debug("userId is empty")
		return false, errors.New("userId is empty")
	}
	if SqlDb == nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//删除数据
	_, err := SqlDb.Exec("delete from tbl_fas_user_oauths where userId=?", userId)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}
