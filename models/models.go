package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego"
	// "GetDomainWhois/whoisonline"
	"fmt"
	"time"
	// "unicode/utf8"
)

type DomainWhois struct {
	Id         int64
	Domain     string    `orm:"index;size(50);null"` //orm 默认为255
	Ip         string    `orm:"null"`
	RegName    string    `orm:"index;null"`
	RegEmail   string    `orm:"index;null"`
	RegPhone   string    `orm:"index;null"`
	Details    string    `orm:"size(10000);null"`
	UpdateTime time.Time `orm:"index;null"`
	QueryTimes int64     `orm:"null"` //查询次数
}

func init() {
	orm.RegisterModel(new(DomainWhois)) //注册模型
}

func GetAllWhois() ([]*DomainWhois, error) {
	o := orm.NewOrm()
	dw := make([]*DomainWhois, 0)
	qs := o.QueryTable("domain_whois")
	// qs.Limit(2)
	// qs.OrderBy("-id")
	// _, err := qs.All(&dw)
	_, err := qs.OrderBy("-id").Limit(20).All(&dw) //返回limite个数的最近查询数据，通过倒序id
	return dw, err

}

//查询用户输入的域名，并且返回到网页上。方法包括数据库已有和未有且通过在线查询的。
func QueryDomain(queryDomain string) DomainWhois {
	o := orm.NewOrm()
	domain := DomainWhois{Domain: queryDomain}
	err := o.Read(&domain, "Domain") //指定字段进行查询
	if err == orm.ErrNoRows {
		fmt.Println("数据库未有该项,正在在线查询")
		domain = QueryOnline(queryDomain)
		o.Insert(&domain)
		return domain
	} else {
		return domain
	}
}

//在线查询得到的域名whois信息
func QueryOnline(queryDomain string) DomainWhois {
	domain := &Domain{}
	domain.ReturnWhois(queryDomain)

	whois := &DomainWhois{
		Domain:     domain.DomainName,
		Ip:         domain.Ip,
		RegName:    string(domain.RegName),
		RegEmail:   domain.RegEmail,
		RegPhone:   domain.RegPhone,
		Details:    string(domain.Details),
		UpdateTime: time.Now(),
	}
	return *whois
}
