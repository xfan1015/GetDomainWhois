package controllers

import (
	// "fmt"
	"GetDomainWhois/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
)

type SrvController struct {
	beego.Controller
}

func (c *SrvController) Get() {
	c.Data["IsSrv"] = true
	c.TplNames = "srv.html"
	var err error
	c.Data["Srv"], c.Data["Count"], err = models.GetAllSrv()
	if err != nil {
		beego.Error(err)
	}

}
