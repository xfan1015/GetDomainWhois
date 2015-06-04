package controllers

import (
	"github.com/astaxie/beego"
)

type WhoisMoreController struct {
	beego.Controller
}

func (c *WhoisMoreController) Get() {
	c.TplNames = "index.html"
}
