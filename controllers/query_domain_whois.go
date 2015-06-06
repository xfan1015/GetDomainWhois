package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
)

type QueryController struct {
	beego.Controller
}

func (c *QueryController) Get() {
	c.Data["IsQuery"] = true
	c.TplNames = "query_domain_whois.html"

}
