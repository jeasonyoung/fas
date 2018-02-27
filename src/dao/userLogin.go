package dao

import (
	"errors"

	"go.uber.org/zap"

	"github.com/satori/go.uuid"

	"fas/src/log"
)

//用户登录流水
type UserLogin struct {
	Id string //登录ID
	UserId string//用户ID
	ChannelId string//所属渠道ID
	Method int8//登录方式(0:本地登录,1:微信,2:支付宝)

	Token string//登录令牌
	IpAddr string//IP地址
	Mac string//设备标识

	ExpiredTime int64//过期时间戳
	Status int8//状态(1:有效,0:无效)
}

//检查ID是否存在
func (u *UserLogin) hasById(id string)(bool, error){
	log.Logger.Debug("hasById", zap.String("id", id))
	if len(id) == 0 {
		log.Logger.Fatal("id is empty")
		return false, errors.New("id is empty")
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	var ret bool
	err := SqlDb.QueryRow("select count(0) > 0 from tbl_fas_user_logins where id=?", id).Scan(&ret)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return ret, nil
}

//根据令牌加载数据
func (u *UserLogin) loadByToken(token string)(bool, error){
	log.Logger.Debug("loadByToken", zap.String("token", token))
	if len(token) == 0 {
		log.Logger.Fatal("token is empty")
		return false, errors.New("token is empty")
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	err := SqlDb.QueryRow("select id,userId,channelId,method,token,ipAddr,mac,expiredTime,status from tbl_fas_user_logins where token=?", token).Scan(
		u.Id, u.UserId, u.ChannelId, u.Method, u.Token, u.IpAddr, u.Mac, u.ExpiredTime, u.Status)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//新增或更新
func (u *UserLogin) SaveOrUpdate() (bool, error){
	log.Logger.Debug("saveOrUpdate")
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	if len(u.Id) == 0 {
		u.Id = uuid.NewV4().String()
	}
	ret,err := u.hasById(u.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	if !ret {//新增
		_, err := SqlDb.Exec("insert into tbl_fas_user_logins(id,userId,channelId,method,token,ipAddr,mac,expiredTime,status) values(?,?,?,?,?,?,?,?,?)",
			u.Id, u.UserId, u.ChannelId, u.Method, u.Token, u.IpAddr, u.Mac, u.ExpiredTime, u.Status)
		if err != nil {
			log.Logger.Fatal(err.Error())
			return false, err
		}
		return true, nil
	}
	//更新
	_, err = SqlDb.Exec("update tbl_fas_user_logins set userId=?,channelId=?,method=?,token=?,ipAddr=?,mac=?,expiredTime=?,status=? where id=?",
		u.UserId, u.ChannelId, u.Method, u.Token, u.IpAddr, u.Mac, u.ExpiredTime, u.Status, u.Id)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

//更新状态
func (u *UserLogin) updateStatus(token string,status bool)(bool, error){
	log.Logger.Debug("updateStatus", zap.String("token", token), zap.Bool("status", status))
	if len(token) == 0 {
		log.Logger.Fatal("token is empty")
		return false, errors.New("token is empty")
	}
	if SqlDb != nil {
		log.Logger.Fatal("sql db is null")
		return false, errors.New("sql db is null")
	}
	//
	s := 0
	if status {
		s = 1
	}
	//
	_, err := SqlDb.Exec("update tbl_fas_user_logins set status=? where token=?", s, token)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return false, err
	}
	return true, nil
}