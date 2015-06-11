package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
)

type PlanController struct {
	beego.Controller
}

func (c *PlanController) Get() {
	c.Data["IsPlan"] = true
	c.TplNames = "plan.html"

}
