package routers

import (
	"GetDomainWhois/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/whois", &controllers.WhoisController{})
	beego.AutoRouter(&controllers.WhoisController{})
	beego.Router("/plan", &controllers.PlanController{})
	beego.Router("/srv", &controllers.SrvController{})
}
