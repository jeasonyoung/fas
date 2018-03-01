package dao

import (
	"errors"

	"github.com/satori/go.uuid"

	"go.uber.org/zap"

	"fas/src/log"
	db "fas/src/database"
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
	log.GetLogInstance().Debug("HasByTypeAndCode", log.Data("type", t), log.Data("code", code))
	if len(code) == 0 {
		log.GetLogInstance().Debug("code is empty")
		return false, errors.New("code is empty")
	}
	if db.SqlDb == nil {
		log.GetLogInstance().Debug("sql db is null")
		return false, errors.New("sql db is null")
	}
	//
	var result bool
	err := db.SqlDb.QueryRow("select count(0) > 0 from tbl_fas_user_oauths where type=? and authCode=?", t, code).Scan(&result)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	return result, nil
}

//新增或更新数据
func (o * OAuth) SaveOrUpdate() (bool, error) {
	log.GetLogInstance().Debug("SaveOrUpdate")
	if db.SqlDb == nil {
		log.GetLogInstance().Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(o.Id) == 0 {
		o.Id = uuid.NewV4().String()
	}
	//检查是否有重复
	ret, err := hasByTypeAndCode(o.Type, o.AuthCode)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	if ret {//
		return false, errors.New("授权码已存在")
	}
	//新增
	_, err = db.SqlDb.Exec("insert into tbl_fas_user_oauths(id,userId,type,authCode) values(?,?,?,?)", o.Id, o.UserId, o.Type, o.AuthCode)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//根据类型和授权码
func (o *OAuth) RemoveByTypeAndCode(t int8, code string)(bool, error){
	log.GetLogInstance().Debug("RemoveByTypeAndCode", log.Data("type", t), log.Data("authCode", code))
	if len(code) == 0 {
		log.GetLogInstance().Fatal("authCode is empty")
		return false, errors.New("authCode is empty")
	}
	if db.SqlDb == nil {
		log.GetLogInstance().Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//删除数据
	_, err := db.SqlDb.Exec("delete from tbl_fas_user_oauths where type=? and authCode=?", t, code)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//根据用户删除数据
func (o *OAuth) RemoveByUser(userId string)(bool, error){
	log.GetLogInstance().Debug("RemoveByUser", log.Data("userId", userId))
	if len(userId) == 0 {
		log.GetLogInstance().Debug("userId is empty")
		return false, errors.New("userId is empty")
	}
	if db.SqlDb == nil {
		log.GetLogInstance().Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//删除数据
	_, err := db.SqlDb.Exec("delete from tbl_fas_user_oauths where userId=?", userId)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	return true, nil
}
