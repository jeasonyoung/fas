package dao

import (
	"errors"

	"github.com/satori/go.uuid"

	"fas/src/log"
	db "fas/src/database"
)

//终端日志
type ClientLog struct {
	Id string //日志ID
	Mac string //设备标识
	UserId string //用户ID

	Type int8 //类型(0:normal,1:warn,2:error)
	Path string//日志文件路径
}

//检查是否存在
func (c *ClientLog) hasById(id string)(bool, error){
	log.GetLogInstance().Debug("hasById", log.Data("id", id))
	if len(id) == 0 {
		return false, errors.New("id is empty")
	}
	if db.SqlDb == nil {
		return false, errors.New("sqlDb is null")
	}
	var state bool
	err := db.SqlDb.QueryRow("select count(0) > 0 from tbl_fas_client_logs where id=?", id).Scan(&state)
	if err != nil {
		return false, err
	}
	return state, err
}

//添加数据
func (c *ClientLog) AddOrUpdate()(bool, error){
	log.GetLogInstance().Debug("add", log.Data("mac", c.Mac), log.Data("userId",c.UserId), log.Data("type", c.Type), log.Data("path", c.Path))
	if len(c.Mac) == 0 || len(c.UserId) == 0 || len(c.Path) == 0 {
		return false, errors.New("mac or userId or path is null")
	}
	if db.SqlDb == nil {
		return false, errors.New("sqlDb is null")
	}
	//新增
	if len(c.Id) == 0 {
		c.Id = uuid.NewV4().String()
		_, err := db.SqlDb.Exec("insert into tbl_fas_client_logs(id,mac,userId,type,path) values(?,?,?,?,?)", c.Id, c.Mac, c.UserId, c.Type, c.Path)
		if err != nil {
			log.GetLogInstance().Fatal(err.Error())
			return false, err
		}
		return true,nil
	}
	//更新
	_, err := db.SqlDb.Exec("update tbl_fas_client_logs set mac=?,userId=?,type=?,path=? where id=?", c.Mac, c.UserId, c.Type, c.Path, c.Id)
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	//
	return true, nil
}