package controllers

import "github.com/astaxie/beego"

//认证控制器
type AuthenController struct {
	beego.Controller
}

//POST 请求
func (a *AuthenController) Post(){
	//a.Ctx.Input.RequestBody
}