package database

import (
	"errors"

	"go.uber.org/zap"
	"fas/src/log"
)

//渠道
type Channel struct {
	Id   string //渠道ID
	Code int    //渠道代码
	Name string //渠道名称
	Status int  //状态(1:正常,0:停用)
}

//根据渠道代码加载数据
func (c *Channel) LoadByCode(code int)(bool, error){
	log.Logger.Debug("loadByCode", zap.Int("code", code))
	if SqlDb == nil {
		return false, errors.New("sqldb is null")
	}
	//查询数据
	err := SqlDb.QueryRow("select id,code,name,status from tbl_fas_sys_channels where code=?", code).Scan(
		c.Id, c.Code, c.Name, c.Status)
	//
	if err != nil {
		log.Logger.Error(err.Error())
		return false, err
	}
	return true, nil
}

//更新状态
func (c *Channel) UpdateStatus(status int)(bool, error){
	log.Logger.Debug("updateStatus", zap.Int("status", status))
	if status < 0 {
		return false, errors.New("status < 0")
	}
	//更新数据
	rs,err := SqlDb.Exec("update tbl_fas_sys_channels set status=? where id=?", status, c.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}