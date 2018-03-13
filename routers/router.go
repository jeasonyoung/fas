// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"

	"fas/controllers"
)

func init() {
	//路由
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			//用户注册
			beego.NSRouter("/register", &controllers.RegisterController{}),
			//用户登录
			beego.NSRouter("/sign", &controllers.AuthenController{}),
			//加载账本
			beego.NSNamespace("/account",
				//用户
				beego.NSRouter("/all", &controllers.AccountController{},"post:LoadAllByUser"),
			),
		),
	)
	//
	beego.AddNamespace(ns)
}
