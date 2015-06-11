package controllers

import (
	// "fmt"
	"GetDomainWhois/models"
	"GetDomainWhois/whoisonline"
	"github.com/astaxie/beego"
	"strings"
)

type WhoisController struct {
	beego.Controller
}

func (c *WhoisController) Get() {

	c.Data["IsQuery"] = true
	c.TplNames = "queryDomain.html"

}

func (c *WhoisController) Query() {
	c.TplNames = "queryDomain.html"
	c.Data["IsQuery"] = true
	queryDomain := c.Input().Get("queryDomain")
	queryDomain = strings.TrimSpace(queryDomain)             //去掉域名前后的空格
	_, _, queryDomain = whoisonline.Extract_tld(queryDomain) //提取用户输入的域名，例如输入www.baidu.com,提取出baidu.com
	if len(queryDomain) == 0 {
		beego.Error("查询不能为空")
		return
		// r, _ := whoisonline.GetDomainWhois("whois.verisign-grs.com", "baidu.com")
		// c.Ctx.WriteString(r)
		// return

	}
	c.Data["domain"] = models.QueryDomain(queryDomain)

}
