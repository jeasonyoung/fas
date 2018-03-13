package controllers

import (
	"fas/models"
)

//账本控制器
type AccountController struct {
	baseController
}

//加载用户账本
func (ac *AccountController) LoadAllByUser(){

}

//全部账本-请求报文体
type ReqAccountQuery struct {
	UserId string `json:"userId"`//用户ID
	Type models.AccountType `json:"type"`//账号类型
	Status models.AccountStatus `json:"status"`//账号状态
}

//全部账本-响应账本数据项
type RespAccountItem struct {
	Id   string `json:"id"`//账本ID
	Name string `json:"name"`//账本名称
	Abbr string `json:"abbr"`//账本简称
	Type models.AccountType `json:"type"`//账本类型
	Status models.AccountStatus `json:"status"`//账本状态
}

