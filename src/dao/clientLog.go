package dao

import (
	"errors"

	"go.uber.org/zap"
	"github.com/satori/go.uuid"

	"fas/src/log"
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
	log.Logger.Debug("hasById", zap.String("id", id))
	if len(id) == 0 {
		return false, errors.New("id is empty")
	}
	if SqlDb == nil {
		return false, errors.New("sqlDb is null")
	}
	var state bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_client_logs where id=?", id).Scan(&state)
	if err != nil {
		return false, err
	}
	return state, err
}

//添加数据
func (c *ClientLog) AddOrUpdate()(bool, error){
	log.Logger.Debug("add", zap.String("mac", c.Mac), zap.String("userId",c.UserId), zap.Int8("type", c.Type), zap.String("path", c.Path))
	if len(c.Mac) == 0 || len(c.UserId) == 0 || len(c.Path) == 0 {
		return false, errors.New("mac or userId or path is null")
	}
	if SqlDb == nil {
		return false, errors.New("sqlDb is null")
	}
	//新增
	if len(c.Id) == 0 {
		c.Id = uuid.NewV4().String()
		_, err := SqlDb.Exec("insert into tbl_fas_client_logs(id,mac,userId,type,path) values(?,?,?,?,?)", c.Id, c.Mac, c.UserId, c.Type, c.Path)
		if err != nil {
			log.Logger.Fatal(err.Error())
			return false, err
		}
		return true,nil
	}
	//更新
	_, err := SqlDb.Exec("update tbl_fas_client_logs set mac=?,userId=?,type=?,path=? where id=?", c.Mac, c.UserId, c.Type, c.Path, c.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	//
	return true, nil
}