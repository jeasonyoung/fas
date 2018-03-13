package net

//查询结果-响应报文体
type RespQueryResultBase struct {
	Total uint `json:"total"`//数据总行数
	Rows interface{} `json:"rows"`//数据
}

//初始化结果对象
func NewRespQueryResultBase(total uint, rows interface{}) *RespQueryResultBase {
	return &RespQueryResultBase{ Total:total, Rows:rows }
}