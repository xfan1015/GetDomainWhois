package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego"
	// "fmt"
)

type DomainWhois struct {
	Id             int64
	Domain         string `orm:"index"` //orm 默认为255
	Ip             string
	TopWhoisServer string
	SecWhoisServer string
	RegName        string `orm:"index"`
	RegEmail       string `orm:"index"`
	RegPhone       string `orm:"index"`
	Details        string `orm:"size(10000)"`
}

func init() {
	orm.RegisterModel(new(DomainWhois)) //注册模型
}

func GetAllWhois() ([]*DomainWhois, error) {
	o := orm.NewOrm()
	dw := make([]*DomainWhois, 0)
	qs := o.QueryTable("domain_whois")
	_, err := qs.All(&dw)
	// fmt.Println(dw)
	return dw, err

}

func QueryDomain(query_domain string) DomainWhois {
	o := orm.NewOrm()
	domain := DomainWhois{Domain: query_domain}
	o.Read(&domain, "Domain") //指定字段进行查询
	return domain
}
