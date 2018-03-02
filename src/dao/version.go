package dao

import (
	"time"
	"errors"

	"fas/src/log"
	db "fas/src/database"
)

//版本管理
type Version struct {
	Id       string //版本ID
	Name     string //版本名称
	Version  uint8 //版本号
	Description string//描述
	ChannelId string //渠道ID

	CheckCode string //校验码
	Status uint8 //状态(1:有效,0:无效)

	StartTime time.Time//生效时间
	Url string //下载地址
}

//加载渠道最新数据
func (v *Version) loadByChannel(channelId string)(bool,error){
	log.GetLogInstance().Debug("loadByChannel", log.Data("channelId", channelId))
	if len(channelId) <= 0 {
		return false, errors.New("channelId is empty")
	}
	if db.SqlDb == nil {
		return false, errors.New("sqldb is null")
	}
	//查询数据
	err := db.SqlDb.QueryRow("select id,name,version,checkCode,status,startTime,url,description,channelId from tbl_fas_sys_versions where channelId=? and status=1 and startTime > now() order by version limit 0,1", channelId).Scan(
		v.Id, v.Name, v.Version, v.CheckCode, v.Status, v.StartTime, v.Url, v.Description, v.ChannelId)
	//
	if err != nil {
		log.GetLogInstance().Fatal(err.Error())
		return false, err
	}
	return true, nil
}
