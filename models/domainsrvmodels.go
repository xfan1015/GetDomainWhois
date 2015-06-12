package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego"
	// "GetDomainWhois/whoisonline"
	// "fmt"
	// "time"
)

type DomainSrv struct {
	Id          int64
	Name        string `orm:"index"` //orm 默认为255
	Tld         string
	RegNameAvl  bool //注册人姓名是否可查
	RegEmailAvl bool //注册人邮箱是否可查
	RegPhoneAvl bool //注册人电话是否可查
}

func init() {
	orm.RegisterModel(new(DomainSrv)) //注册模型
}
