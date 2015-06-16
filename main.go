package main

import (
	_ "GetDomainWhois/models"
	_ "GetDomainWhois/routers"
	//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)                                                                                    //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", "root:cynztt@/malicious_detect?charset=utf8&loc="+url.QueryEscape("Asia/Shanghai")) //注册数据库,务必加上时区设置，否则出错
}

func main() {
	orm.Debug = true
	//	fmt.Println("test")
	orm.RunSyncdb("default", false, true) //false的意思是不用每次重新删掉重建
	beego.Run()
}
