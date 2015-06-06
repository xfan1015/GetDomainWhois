package routers

import (
	"GetDomainWhois/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/whoismoreinfo", &controllers.WhoisMoreController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/query", &controllers.QueryController{})
}
