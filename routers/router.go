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
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/sign",
			beego.NSInclude(
				&controllers.AuthenController{},
			),
		),
	)
	//
	beego.AddNamespace(ns)
}
