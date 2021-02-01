// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"BeeApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/addEvent",
			beego.NSInclude(
				&controllers.EventItemCtrl{},
			),
		),
		//beego.NSRouter("/student", &controllers.StudentController{}, "get:Get"),
		//beego.NSRouter("/addEvent",&controllers.EventItemCtrl{},"Post:AddEvent"),
	)
	beego.AddNamespace(ns)
}
