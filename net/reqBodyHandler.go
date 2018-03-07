package net

//请求报文处理方法接口
type ReqBodyHandler interface {

	//校验数据
	Valid() error

	//保存数据
	Save() error
}
