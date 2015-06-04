package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	c.TplNames = "login.html"

}
