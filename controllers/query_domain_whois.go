package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
	"GetDomainWhois/models"
)

type QueryController struct {
	beego.Controller
}

func (c *QueryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "query":
		query_domain := c.Input().Get("query_domain")
		if len(query_domain) == 0 {
			break
		}
		c.Data["domain"] = models.QueryDomain(query_domain)
		// c.Redirect("/query", 301)
		// return

	}

	c.Data["IsQuery"] = true
	c.TplNames = "query_domain_whois.html"

}
