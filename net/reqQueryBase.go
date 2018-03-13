package net

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

//分页查询-请求报文体基类
type ReqQueryBase struct {
	//Sort  string `json:"sort"`//排序字段
	//Order string `json:"order"`//排序方向
	Rows  uint `json:"rows"`//每页数据
	Index uint `json:"index"`//页码
	Query interface{} `json:"query"`//查询条件
}

//查询类型转换
func (q *ReqQueryBase) ToQueryTarget(target interface{}) error {
	if q.Query != nil && target != nil {
		//将body对象json
		data, err := json.Marshal(q.Query)
		if err != nil {
			logs.Debug("ToQueryTarget[%v]-exp:%v", q.Query, err.Error())
			return err
		}
		//json解析
		if len(data) > 0 {
			return json.Unmarshal(data, target)
		}
	}
	return nil
}