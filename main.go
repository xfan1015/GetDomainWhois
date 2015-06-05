package main

import (
	_ "GetDomainWhois/models"
	_ "GetDomainWhois/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)                                              //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", "root:cynztt@/malicious_detect?charset=utf8") //注册数据库
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true) //false的意思是不用每次重新删掉重建
	beego.Run()
}
