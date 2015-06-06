package controllers

import (
	"GetDomainWhois/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	var err error
	c.Data["Domainwhois"], err = models.GetAllWhois()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "index.html"
}
